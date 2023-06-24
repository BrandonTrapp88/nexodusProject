package handlers

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm/clause"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nexodus-io/nexodus/internal/models"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

func (api *API) metadataForDevice(c *gin.Context, deviceId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("device_id = ?", deviceId)
	}
}

// GetDeviceMetadata lists metadata for a device
// @Summary      Get Device Metadata
// @Id  		 GetDeviceMetadata
// @Tags         Devices
// @Description  Lists metadata for a device
// @Param        id   path      string  true "Device ID"
// @Accept	     json
// @Produce      json
// @Success      200  {object}  models.DeviceMetadata
// @Failure      500  {object}  models.BaseError
// @Router       /api/devices/{id}/metadata [get]
func (api *API) GetDeviceMetadata(c *gin.Context) {
	ctx, span := tracer.Start(c.Request.Context(), "GetDeviceMetadata", trace.WithAttributes(
		attribute.String("id", c.Param("id")),
	))
	defer span.End()
	deviceId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewBadPathParameterError("id"))
		return
	}

	var metadataInstances []models.DeviceMetadataInstance

	err = api.transaction(ctx, func(tx *gorm.DB) error {
		var device models.Device
		result := api.db.WithContext(ctx).
			Scopes(api.DeviceIsOwnedByCurrentUser(c)).
			First(&device, "id = ?", deviceId)
		if result.Error != nil {
			return result.Error
		}

		result = api.db.WithContext(ctx).
			Scopes(
				api.metadataForDevice(c, deviceId.String()),
				FilterAndPaginate(&models.DeviceMetadataInstance{}, c, "key"),
			).
			Find(&metadataInstances)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching keys from db"})
			return result.Error
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		api.logger.Errorf("error fetching metadata: %s", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	result := models.DeviceMetadata{
		DeviceID: deviceId.String(),
		Metadata: make(map[string]models.DeviceMetadataValue),
	}
	for _, metadata := range metadataInstances {
		result.Metadata[metadata.Key] = models.DeviceMetadataValue{
			Value:    metadata.Value,
			Revision: metadata.Revision,
		}
	}
	c.JSON(http.StatusOK, result)
}

// GetDeviceMetadata Get value for a metadata key on a device
// @Summary      Get Device Metadata
// @Id  		 GetDeviceMetadataKey
// @Tags         Devices
// @Description  Get metadata for a device
// @Param        id   path      string  true "Device ID"
// @Param        key  path      string  true "Metadata Key"
// @Accept	     json
// @Produce      json
// @Success      200  {object}  models.DeviceMetadataValue
// @Failure      501  {object}  models.BaseError
// @Router       /api/devices/{id}/metadata/{key} [get]
func (api *API) GetDeviceMetadataKey(c *gin.Context) {

	ctx, span := tracer.Start(c.Request.Context(), "GetDeviceMetadataKey", trace.WithAttributes(
		attribute.String("id", c.Param("id")),
		attribute.String("key", c.Param("key")),
	))
	defer span.End()
	deviceId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewBadPathParameterError("id"))
		return
	}
	key := c.Param("key")

	var metadataInstance models.DeviceMetadataInstance
	err = api.transaction(ctx, func(tx *gorm.DB) error {
		var device models.Device
		result := api.db.WithContext(ctx).
			Scopes(api.DeviceIsOwnedByCurrentUser(c)).
			First(&device, "id = ?", deviceId)
		if result.Error != nil {
			return result.Error
		}

		result = api.db.WithContext(ctx).
			Scopes(api.metadataForDevice(c, deviceId.String())).
			Where("key = ?", key).
			First(&metadataInstance)

		return result.Error
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		api.logger.Errorf("error fetching metadata: %s", err)
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, models.DeviceMetadataValue{
		Value:    metadataInstance.Value,
		Revision: metadataInstance.Revision,
	})
}

// UpdateDeviceMetadataKey Set value for a metadata key on a device
// @Summary      Set Device Metadata by key
// @Id  		 UpdateDeviceMetadataKey
// @Tags         Devices
// @Description  Set metadata key for a device
// @Param        id    path      string          true  "Device ID"
// @Param        key   path      string          false "Metadata Key"
// @Param		 value body      any true "Metadata Value"
// @Accept	     json
// @Produce      json
// @Success      200  {object}  models.DeviceMetadataValue
// @Failure      501  {object}  models.BaseError
// @Router       /api/devices/{id}/metadata/{key} [put]
func (api *API) UpdateDeviceMetadataKey(c *gin.Context) {

	ctx, span := tracer.Start(c.Request.Context(), "GetDeviceMetadataKey", trace.WithAttributes(
		attribute.String("id", c.Param("id")),
		attribute.String("key", c.Param("key")),
	))
	defer span.End()
	deviceId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewBadPathParameterError("id"))
		return
	}
	key := c.Param("key")

	var request json.RawMessage
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.NewBadPayloadError())
		return
	}

	metadataInstance := models.DeviceMetadataInstance{
		DeviceID: deviceId,
		Key:      key,
		Value:    request,
	}
	err = api.transaction(ctx, func(tx *gorm.DB) error {
		var device models.Device
		result := api.db.WithContext(ctx).
			Scopes(api.DeviceIsOwnedByCurrentUser(c)).
			First(&device, "id = ?", deviceId)
		if result.Error != nil {
			return result.Error
		}

		result = tx.
			Clauses(clause.Returning{Columns: []clause.Column{{Name: "revision"}}}).
			Save(&metadataInstance)
		return result.Error
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		api.logger.Errorf("error updating metadata: %s", err)
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, models.DeviceMetadataValue{
		Value:    metadataInstance.Value,
		Revision: metadataInstance.Revision,
	})

}

// DeleteDeviceMetadata Delete all metadata or a specific key on a device
// @Summary      Delete all Device metadata
// @Id  		 DeleteDeviceMetadata
// @Tags         Devices
// @Description  Delete all metadata for a device
// @Param        id   path      string  true "Device ID"
// @Success      204
// @Failure      501  {object}  models.BaseError
// @Router       /api/devices/{id}/metadata [delete]
func (api *API) DeleteDeviceMetadata(c *gin.Context) {

	ctx, span := tracer.Start(c.Request.Context(), "GetDeviceMetadataKey", trace.WithAttributes(
		attribute.String("id", c.Param("id")),
	))
	defer span.End()
	deviceId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewBadPathParameterError("id"))
		return
	}

	err = api.transaction(ctx, func(tx *gorm.DB) error {
		var device models.Device
		result := api.db.WithContext(ctx).
			Scopes(api.DeviceIsOwnedByCurrentUser(c)).
			First(&device, "id = ?", deviceId)
		if result.Error != nil {
			return result.Error
		}

		result = tx.Delete(&models.DeviceMetadataInstance{}, "device_id", deviceId)
		return result.Error
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		api.logger.Errorf("error deleting metadata: %s", err)
		c.JSON(http.StatusInternalServerError, err)
	}

	c.Status(http.StatusNoContent)
}

// DeleteDeviceMetadataKey Delete all metadata or a specific key on a device
// @Summary      Delete a Device metadata key
// @Id  		 DeleteDeviceMetadataKey
// @Tags         Devices
// @Description  Delete a metadata key for a device
// @Param        id   path      string  true "Device ID"
// @Param        key  path      string  false "Metadata Key"
// @Success      204
// @Failure      501  {object}  models.BaseError
// @Router       /api/devices/{id}/metadata/{key} [delete]
func (api *API) DeleteDeviceMetadataKey(c *gin.Context) {
	ctx, span := tracer.Start(c.Request.Context(), "GetDeviceMetadataKey", trace.WithAttributes(
		attribute.String("id", c.Param("id")),
		attribute.String("key", c.Param("key")),
	))
	defer span.End()
	deviceId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewBadPathParameterError("id"))
		return
	}
	key := c.Param("key")

	err = api.transaction(ctx, func(tx *gorm.DB) error {
		var device models.Device
		result := api.db.WithContext(ctx).
			Scopes(api.DeviceIsOwnedByCurrentUser(c)).
			First(&device, "id = ?", deviceId)
		if result.Error != nil {
			return result.Error
		}

		result = tx.Delete(&models.DeviceMetadataInstance{
			DeviceID: deviceId,
			Key:      key,
		})
		return result.Error
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		api.logger.Errorf("error deleting metadata: %s", err)
		c.JSON(http.StatusInternalServerError, err)
	}

	c.Status(http.StatusNoContent)
}
