package service

import (
	"fmt"

	"github.com/mattli001/AI-100x-SE-Join-Quest/src/domain"
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
