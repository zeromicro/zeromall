// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type LicenseCreateReq struct {
	ConsumerID string `json:"consumer_id"`
	ProductID  string `json:"product_id"`
	ShopID     string `json:"shop_id"`
}

type LicenseCreateResp struct {
	ConsumerID string `json:"consumer_id"`
	ProductID  string `json:"product_id"`
	ShopID     string `json:"shop_id"`
	PublicKey  string `json:"public_key"`
	SecretKey  string `json:"secret_key"`
}

type LicenseGetReq struct {
	PublicKey string `json:"public_key"`
}

type LicenseGetResp struct {
	PublicKey       string `json:"public_key"`
	SecretKey       string `json:"secret_key"`
	ConsumerID      string `json:"consumer_id"`
	ProductID       string `json:"product_id"`
	ShopID          string `json:"shop_id"`
	Status          string `json:"status"`
	DeviceLimit     string `json:"device_limit"`
	ActivatedDevice int    `json:"activated_device"`
}
