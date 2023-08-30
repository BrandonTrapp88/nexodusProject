package stun

import (
	"fmt"
	"github.com/libp2p/go-reuseport"
	"github.com/pion/stun"
	"go.uber.org/zap"
	"net/netip"
	"runtime"
)

func RequestWithReusePort(logger *zap.SugaredLogger, stunServer string, srcPort int) (netip.AddrPort, error) {
	logger.Debugf("dialing stun Server %s", stunServer)
	conn, err := reuseport.Dial("udp4", fmt.Sprintf(":%d", srcPort), stunServer)
	if err != nil {
		// Windows is currently not capable of binding to the source wg port to source STUN requests
		if runtime.GOOS != "windows" {
			logger.Errorf("stun dialing timed out %v", err)
		}
		return netip.AddrPort{}, fmt.Errorf("failed to dial stun Server %s: %w", stunServer, err)
	}
	defer func() {
		_ = conn.Close()
	}()

	c, err := stun.NewClient(conn)
	if err != nil {
		logger.Error(err)
		return netip.AddrPort{}, err
	}
	defer func() {
		_ = c.Close()
	}()

	// Building binding request with random transaction id.
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	// Sending request to STUN Server, waiting for response message.
	var xorAddr stun.XORMappedAddress
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			if res.Error.Error() == "transaction is timed out" {
				logger.Debugf("STUN transaction timed out, if this continues check if a firewall is blocking UDP connections to %s", stunServer)
			} else {
				logger.Debug(res.Error)
			}
			return
		}
		// Decoding XOR-MAPPED-ADDRESS attribute from message.
		if err := xorAddr.GetFrom(res.Message); err != nil {
			return
		}
	}); err != nil {
		return netip.AddrPort{}, err
	}

	xorBinding, err := netip.ParseAddrPort(xorAddr.String())
	if err != nil {
		return netip.AddrPort{}, fmt.Errorf("failed to parse a valid address:port binding from the stun response: %w", err)
	}
	logger.Debugf("reflexive binding is: %s", xorBinding.String())

	return xorBinding, nil
}
