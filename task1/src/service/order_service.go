package service

import (
	"fmt"

	"github.com/mattli001/AI-100x-SE-Join-Quest/task1/src/domain"
)

// OrderService 處理訂單相關業務邏輯
type OrderService struct {
	promotions []domain.Promotion
}

// NewOrderService 建立新的 OrderService
func NewOrderService() *OrderService {
	return &OrderService{
		promotions: make([]domain.Promotion, 0),
	}
}

// AddPromotion 增加優惠活動
func (s *OrderService) AddPromotion(promotion domain.Promotion) {
	s.promotions = append(s.promotions, promotion)
}

// ClearPromotions 清除所有優惠活動
func (s *OrderService) ClearPromotions() {
	s.promotions = make([]domain.Promotion, 0)
}

// ProcessOrder 處理訂單並套用優惠
func (s *OrderService) ProcessOrder(items []domain.OrderItem) (*domain.Order, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("no items in order")
	}

	order := &domain.Order{
		Items:      items,
		BonusItems: make([]domain.OrderItem, 0),
	}

	// 計算原始金額
	originalAmount := 0
	for _, item := range items {
		originalAmount += item.Product.Price * item.Quantity
	}
	order.OriginalAmount = originalAmount

	// 套用優惠
	order.Discount = 0
	for _, promotion := range s.promotions {
		err := promotion.Apply(order)
		if err != nil {
			return nil, err
		}
	}

	// 計算最終金額
	order.TotalAmount = order.OriginalAmount - order.Discount

	return order, nil
}

// GetOrderSummary 取得訂單摘要
func (s *OrderService) GetOrderSummary(order *domain.Order) domain.OrderSummary {
	return domain.OrderSummary{
		OriginalAmount: order.OriginalAmount,
		Discount:       order.Discount,
		TotalAmount:    order.TotalAmount,
	}
}

// GetReceivedItems 取得客戶實際收到的商品
func (s *OrderService) GetReceivedItems(order *domain.Order) []domain.ReceivedItem {
	receivedItems := make([]domain.ReceivedItem, 0)
	productQuantityMap := make(map[string]int)

	// 計算購買的商品數量
	for _, item := range order.Items {
		productQuantityMap[item.Product.Name] += item.Quantity
	}

	// 加上贈品數量
	for _, bonusItem := range order.BonusItems {
		productQuantityMap[bonusItem.Product.Name] += bonusItem.Quantity
	}

	// 轉換為 ReceivedItem 格式
	for _, item := range order.Items {
		productName := item.Product.Name
		totalQuantity := productQuantityMap[productName]

		// 檢查是否已經加入過這個商品
		found := false
		for _, received := range receivedItems {
			if received.ProductName == productName {
				found = true
				break
			}
		}

		if !found {
			receivedItem := domain.ReceivedItem{
				ProductName: productName,
				Quantity:    totalQuantity,
			}
			receivedItems = append(receivedItems, receivedItem)
		}
	}

	return receivedItems
}

// GetDiscountDetails 取得折扣詳情
func (s *OrderService) GetDiscountDetails(order *domain.Order) []domain.DiscountDetail {
	details := make([]domain.DiscountDetail, 0)

	// 尋找雙十一促銷
	var doubleElevenPromotion *domain.DoubleElevenPromotion
	for _, promotion := range s.promotions {
		if de, ok := promotion.(*domain.DoubleElevenPromotion); ok {
			doubleElevenPromotion = de
			break
		}
	}

	// 如果沒有雙十一促銷，則所有商品都是原價
	if doubleElevenPromotion == nil {
		return s.getRegularPriceDetails(order)
	}

	// 按商品分組計算詳情
	productQuantityMap := make(map[string]int)
	productPriceMap := make(map[string]int)

	for _, item := range order.Items {
		productName := item.Product.Name
		productQuantityMap[productName] += item.Quantity
		productPriceMap[productName] = item.Product.Price
	}

	// 對每個商品計算折扣詳情
	for _, item := range order.Items {
		productName := item.Product.Name
		totalQuantity := productQuantityMap[productName]
		unitPrice := productPriceMap[productName]

		// 檢查是否已經處理過這個商品
		found := false
		for _, detail := range details {
			if detail.ProductName == productName {
				found = true
				break
			}
		}

		if !found {
			detail := doubleElevenPromotion.CalculateDiscountDetail(productName, totalQuantity, unitPrice)
			details = append(details, detail)
		}
	}

	return details
}

// getRegularPriceDetails 取得原價詳情（無促銷時使用）
func (s *OrderService) getRegularPriceDetails(order *domain.Order) []domain.DiscountDetail {
	details := make([]domain.DiscountDetail, 0)

	// 按商品分組
	productQuantityMap := make(map[string]int)
	productPriceMap := make(map[string]int)

	for _, item := range order.Items {
		productName := item.Product.Name
		productQuantityMap[productName] += item.Quantity
		productPriceMap[productName] = item.Product.Price
	}

	// 對每個商品建立詳情
	for _, item := range order.Items {
		productName := item.Product.Name
		totalQuantity := productQuantityMap[productName]
		unitPrice := productPriceMap[productName]

		// 檢查是否已經處理過這個商品
		found := false
		for _, detail := range details {
			if detail.ProductName == productName {
				found = true
				break
			}
		}

		if !found {
			detail := domain.DiscountDetail{
				ProductName:        productName,
				DiscountedQuantity: 0,
				RegularQuantity:    totalQuantity,
				DiscountedAmount:   0,
				RegularAmount:      totalQuantity * unitPrice,
			}
			details = append(details, detail)
		}
	}

	return details
}
