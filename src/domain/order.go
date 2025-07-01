package domain

// Product 代表商品
type Product struct {
	Name     string
	Category string
	Price    int // 以分為單位
}

// OrderItem 代表訂單項目
type OrderItem struct {
	Product  Product
	Quantity int
}

// Order 代表訂單
type Order struct {
	Items          []OrderItem
	BonusItems     []OrderItem // 贈品項目
	OriginalAmount int
	Discount       int
	TotalAmount    int
}

// OrderSummary 代表訂單摘要
type OrderSummary struct {
	OriginalAmount int
	Discount       int
	TotalAmount    int
}

// ReceivedItem 代表客戶實際收到的商品
type ReceivedItem struct {
	ProductName string
	Quantity    int
}

// Promotion 代表優惠活動介面
type Promotion interface {
	Apply(order *Order) error
}

// ThresholdDiscountPromotion 代表滿額折扣優惠
type ThresholdDiscountPromotion struct {
	Threshold int
	Discount  int
}

// Apply 實作滿額折扣優惠邏輯
func (p *ThresholdDiscountPromotion) Apply(order *Order) error {
	if order.OriginalAmount >= p.Threshold {
		order.Discount += p.Discount
	}
	return nil
}

// BuyOneGetOnePromotion 代表買一送一優惠
type BuyOneGetOnePromotion struct {
	Category string
}

// Apply 實作買一送一優惠邏輯
func (p *BuyOneGetOnePromotion) Apply(order *Order) error {
	bonusMap := make(map[string]*OrderItem)

	for _, item := range order.Items {
		if item.Product.Category == p.Category {
			productName := item.Product.Name

			// 買一送一邏輯：每個符合條件的商品都送一個，不論購買數量
			if _, exists := bonusMap[productName]; exists {
				// 如果已經有這個商品的贈品，不再增加（每個商品只享受一次優惠）
				continue
			} else {
				bonusMap[productName] = &OrderItem{
					Product:  item.Product,
					Quantity: 1, // 每個商品只送一個
				}
			}
		}
	}

	// 將贈品加入訂單
	for _, bonusItem := range bonusMap {
		order.BonusItems = append(order.BonusItems, *bonusItem)
	}

	return nil
}
