package config

type JsonResponseEasyDonate struct {
	Success  bool `json:"success"`
	Response struct {
		URL     string `json:"url"`
		Payment struct {
			ID          int         `json:"id"`
			Customer    string      `json:"customer"`
			Email       interface{} `json:"email"`
			ServerID    int         `json:"server_id"`
			PaymentType bool        `json:"payment_type"`
			ShopID      int         `json:"shop_id"`
			UpdatedAt   string      `json:"updated_at"`
			CreatedAt   string      `json:"created_at"`
			Enrolled    float64     `json:"enrolled"`
			Cost        int         `json:"cost"`
			Server      struct {
				ID           int    `json:"id"`
				Name         string `json:"name"`
				IP           string `json:"ip"`
				Port         string `json:"port"`
				Version      string `json:"version"`
				IsPortHidden int    `json:"is_port_hidden"`
				HideIP       int    `json:"hide_ip"`
				IsHidden     int    `json:"is_hidden"`
				ShopID       int    `json:"shop_id"`
				CreatedAt    string `json:"created_at"`
				UpdatedAt    string `json:"updated_at"`
			} `json:"server"`
			Products []struct {
				ID               int         `json:"id"`
				ProductID        int         `json:"product_id"`
				Name             string      `json:"name"`
				Price            int         `json:"price"`
				OldPrice         interface{} `json:"old_price"`
				Type             string      `json:"type"`
				Number           int         `json:"number"`
				Commands         []string    `json:"commands"`
				AdditionalFields interface{} `json:"additional_fields"`
				Description      interface{} `json:"description"`
				PaymentID        int         `json:"payment_id"`
				Amount           int         `json:"amount"`
				Image            string      `json:"image"`
				FirstDelete      int         `json:"first_delete"`
				CreatedAt        string      `json:"created_at"`
				UpdatedAt        string      `json:"updated_at"`
			} `json:"products"`
		} `json:"payment"`
	} `json:"response"`
}

type InfoRequest struct {
	Type      string `json:"type"`
	Message   string `json:"message"`
	OrderID   int    `json:"orderId"`
	OrderHash string `json:"orderHash"`
	Location  string `json:"location"`
}
