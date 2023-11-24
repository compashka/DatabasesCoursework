package rest

import (
	"app/internal/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{db: db}
}

//Auth

func (h *handler) NewAuth() (jwt.Auth, error) {
	return jwt.New(jwt.Auth{
		SecretKey: []byte("secret"),
		Authenticator: func(c *gin.Context) (jwt.MapClaims, error) {
			var req struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			if err := c.ShouldBind(&req); err != nil {
				return nil, jwt.ErrorAuthenticationFailed
			}

			var user models.User
			result := h.db.Where("username = ?", req.Username).First(&user)
			if result.Error != nil {
				return nil, jwt.ErrorAuthenticationFailed
			}

			if user.Password != req.Password {
				return nil, jwt.ErrorAuthenticationFailed
			}

			return jwt.MapClaims{
				"username": user.Username,
				"role":     user.Role,
			}, nil
		},
		UserFetcher: func(c *gin.Context, claims jwt.MapClaims) (interface{}, error) {
			username, ok := claims["username"].(string)
			if !ok {
				return nil, nil
			}
			var user = models.User{
				Username: username,
			}
			result := h.db.First(&user)
			if result.Error != nil {
				return nil, nil
			}
			return user, nil
		},
	})
}

func (h *handler) Register(c *gin.Context) {
	var user = models.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, err)
		return
	}
	user.Role = "WORKER"
	result := h.db.Create(&user)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, user)
}

func Worker(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims) == models.RoleWorker
	})
}

func Dispatcher(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims) == models.RoleDispatcher
	})
}

func role(claims jwt.MapClaims) models.Role {
	return models.Role(claims["role"].(string))
}

//Get by primary key

func (h *handler) GetSubstationByPK(c *gin.Context) {
	pk := c.Param("pk")
	var obj = models.Substation{
		Name: pk,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetFactoryByPK(c *gin.Context) {
	pk := c.Param("pk")
	var obj = models.Factory{
		Name: pk,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetRangeOfHighVoltageEquipmentByPK(c *gin.Context) {
	pk := c.Param("pk")
	pkInt, _ := strconv.Atoi(pk)
	var obj = models.RangeOfHighVoltageEquipment{
		ID: pkInt,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetCableLineByPK(c *gin.Context) {
	pk := c.Param("pk")
	var obj = models.CableLine{
		Mark: pk,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetTireSectionByPK(c *gin.Context) {
	pk := c.Param("pk")
	var obj = models.TireSection{
		Name: pk,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetCellKVLByPK(c *gin.Context) {
	pk := c.Param("pk")
	var obj = models.CellKVL{
		DispatchName: pk,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetFuseByPK(c *gin.Context) {
	pk := c.Param("pk")
	var obj = models.Fuse{
		Mark: pk,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetCellTNByPK(c *gin.Context) {
	pk := c.Param("pk")
	var obj = models.CellTN{
		DispatchName: pk,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetCellTSNByPK(c *gin.Context) {
	pk := c.Param("pk")
	var obj = models.CellTSN{
		DispatchName: pk,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetNSSByPK(c *gin.Context) {
	pk := c.Param("pk")
	pkInt, _ := strconv.Atoi(pk)
	var obj = models.NSS{
		ID: pkInt,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetRangeOfStandardVoltageByPK(c *gin.Context) {
	pk := c.Param("pk")
	pkInt, _ := strconv.Atoi(pk)
	var obj = models.RangeOfStandardVoltage{
		ID: pkInt,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetTypeOfTransformerByPK(c *gin.Context) {
	pk := c.Param("pk")
	var obj = models.TypeOfTransformer{
		Type: pk,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetTransformerByPK(c *gin.Context) {
	pk := c.Param("pk")
	pkInt, _ := strconv.Atoi(pk)
	var obj = models.Transformer{
		FactoryNumber: pkInt,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetWorkerByPK(c *gin.Context) {
	pk := c.Param("pk")
	pkInt, _ := strconv.Atoi(pk)
	var obj = models.Worker{
		ID: pkInt,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, obj)
}

func (h *handler) GetRequestByPK(c *gin.Context) {
	pk := c.Param("pk")
	pkInt, _ := strconv.Atoi(pk)
	var obj = models.Request{
		ID: pkInt,
	}
	result := h.db.First(&obj)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}

	c.JSON(200, obj)
}

//Get transformers by location of substation

func (h *handler) GetTransformersByLocation(c *gin.Context) {
	location := c.Param("location")
	var substations []models.Substation
	result := h.db.Raw("SELECT * "+
		"FROM transformers "+
		"JOIN substations on transformers.substation = substations.name "+
		"WHERE substations.location = @location",
		sql.Named("location", location)).Find(&substations)
	if result.RowsAffected == 0 {
		c.JSON(400, "record not found")
		return
	}
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, substations)
}

//Get all requests

func (h *handler) GetAllRequests(c *gin.Context) {
	var requests []models.Request
	result := h.db.Find(&requests)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, requests)
}

//Get requests by worker_id

func (h *handler) GetRequestsByWorkerId(c *gin.Context) {
	var requests []models.Request
	workerId := c.Param("workerId")
	result := h.db.Where("worker_id = ?", workerId).Find(&requests)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, requests)
}

//Create request

func (h *handler) CreateRequest(c *gin.Context) {
	transformerId := c.PostForm("transformer_factory_number")
	transformerIdInt, _ := strconv.Atoi(transformerId)
	var transformer = models.Transformer{
		FactoryNumber: transformerIdInt,
	}
	result := h.db.First(&transformer)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}

	workerId := c.PostForm("worker_id")
	workerIdInt, _ := strconv.Atoi(workerId)
	var worker = models.Worker{
		ID: workerIdInt,
	}
	result = h.db.First(&worker)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}

	request := models.Request{
		TransformerFactoryNumber: transformer.FactoryNumber,
		WorkerID:                 worker.ID,
		IsCompleted:              false,
		DateOpened:               time.Now(),
	}

	h.db.Model(&transformer).Association("TransformerFactoryNumber").Append(&request)
	h.db.Model(&worker).Association("WorkerID").Append(&request)
	result = h.db.Create(&request)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, request)
}

//Update request (close it)

func (h *handler) UpdateRequest(c *gin.Context) {
	transformerId := c.PostForm("transformer_factory_number")
	transformerIdInt, _ := strconv.Atoi(transformerId)
	var transformer = models.Transformer{
		FactoryNumber: transformerIdInt,
	}
	result := h.db.First(&models.Transformer{FactoryNumber: transformerIdInt})
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}

	workerId := c.PostForm("worker_id")
	workerIdInt, _ := strconv.Atoi(workerId)
	var worker = models.Worker{
		ID: workerIdInt,
	}
	result = h.db.First(&worker)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}

	request := models.Request{
		TransformerFactoryNumber: transformer.FactoryNumber,
		WorkerID:                 worker.ID,
	}
	h.db.First(&request)
	request.IsCompleted = true
	request.DateClosed = time.Now()

	h.db.Model(&transformer).Association("TransformerFactoryNumber").Append(&request)
	h.db.Model(&worker).Association("WorkerID").Append(&request)
	result = h.db.Save(&request)
	if result.Error != nil {
		c.JSON(400, result.Error.Error())
		return
	}
	c.JSON(200, request)
}
