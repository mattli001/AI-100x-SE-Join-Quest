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

// DiscountDetail 代表折扣詳情
type DiscountDetail struct {
	ProductName        string
	DiscountedQuantity int
	RegularQuantity    int
	DiscountedAmount   int
	RegularAmount      int
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

// DoubleElevenPromotion 代表雙十一促銷優惠
type DoubleElevenPromotion struct {
	ThresholdQuantity int
	DiscountRate      float64
}

// NewDoubleElevenPromotion 建立新的雙十一促銷
func NewDoubleElevenPromotion() *DoubleElevenPromotion {
	return &DoubleElevenPromotion{
		ThresholdQuantity: 10,
		DiscountRate:      0.8,
	}
}

// Apply 實作雙十一促銷優惠邏輯
func (p *DoubleElevenPromotion) Apply(order *Order) error {
	// 按商品分組計算折扣
	productQuantityMap := make(map[string]int)
	productPriceMap := make(map[string]int)

	for _, item := range order.Items {
		productName := item.Product.Name
		productQuantityMap[productName] += item.Quantity
		productPriceMap[productName] = item.Product.Price
	}

	totalDiscount := 0

	// 對每個商品計算折扣
	for productName, totalQuantity := range productQuantityMap {
		if totalQuantity >= p.ThresholdQuantity {
			unitPrice := productPriceMap[productName]
			discount := p.calculateDiscount(totalQuantity, unitPrice)
			totalDiscount += discount
		}
	}

	order.Discount += totalDiscount
	return nil
}

// calculateDiscount 計算單一商品的折扣金額
func (p *DoubleElevenPromotion) calculateDiscount(totalQuantity, unitPrice int) int {
	// 計算可享受折扣的組數
	discountGroups := totalQuantity / p.ThresholdQuantity
	discountQuantity := discountGroups * p.ThresholdQuantity

	// 計算折扣金額
	originalAmount := discountQuantity * unitPrice
	discountedAmount := int(float64(originalAmount) * p.DiscountRate)
	return originalAmount - discountedAmount
}

// CalculateDiscountDetail 計算單一商品的折扣詳情
func (p *DoubleElevenPromotion) CalculateDiscountDetail(productName string, totalQuantity, unitPrice int) DiscountDetail {
	var discountedQuantity, regularQuantity int
	var discountedAmount, regularAmount int

	if totalQuantity >= p.ThresholdQuantity {
		// 計算可享受折扣的組數
		discountGroups := totalQuantity / p.ThresholdQuantity
		discountedQuantity = discountGroups * p.ThresholdQuantity
		regularQuantity = totalQuantity - discountedQuantity

		// 計算金額
		discountedOriginalAmount := discountedQuantity * unitPrice
		discountedAmount = int(float64(discountedOriginalAmount) * p.DiscountRate)
		regularAmount = regularQuantity * unitPrice
	} else {
		// 不符合折扣條件
		discountedQuantity = 0
		regularQuantity = totalQuantity
		discountedAmount = 0
		regularAmount = totalQuantity * unitPrice
	}

	return DiscountDetail{
		ProductName:        productName,
		DiscountedQuantity: discountedQuantity,
		RegularQuantity:    regularQuantity,
		DiscountedAmount:   discountedAmount,
		RegularAmount:      regularAmount,
	}
}
