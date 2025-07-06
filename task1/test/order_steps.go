package test

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cucumber/godog"
	"github.com/mattli001/AI-100x-SE-Join-Quest/task1/src/domain"
	"github.com/mattli001/AI-100x-SE-Join-Quest/task1/src/service"
)

// OrderTestContext 測試上下文
type OrderTestContext struct {
	orderService *service.OrderService
	orderItems   []domain.OrderItem
	order        *domain.Order
	lastError    error
}

// RegisterSteps 註冊所有測試步驟
func (ctx *OrderTestContext) RegisterSteps(sc *godog.ScenarioContext) {
	sc.Before(ctx.beforeScenario)

	// Given steps
	sc.Step(`^no promotions are applied$`, ctx.noPromotionsAreApplied)
	sc.Step(`^the threshold discount promotion is configured:$`, ctx.theThresholdDiscountPromotionIsConfigured)
	sc.Step(`^the buy one get one promotion for cosmetics is active$`, ctx.theBuyOneGetOnePromotionForCosmeticsIsActive)

	// When steps
	sc.Step(`^a customer places an order with:$`, ctx.aCustomerPlacesAnOrderWith)

	// Then steps
	sc.Step(`^the order summary should be:$`, ctx.theOrderSummaryShouldBe)
	sc.Step(`^the customer should receive:$`, ctx.theCustomerShouldReceive)
}

func (ctx *OrderTestContext) beforeScenario(ctx2 context.Context, sc *godog.Scenario) (context.Context, error) {
	ctx.orderService = service.NewOrderService()
	ctx.orderItems = make([]domain.OrderItem, 0)
	ctx.order = nil
	ctx.lastError = nil
	return ctx2, nil
}

func (ctx *OrderTestContext) noPromotionsAreApplied() error {
	// 清除所有優惠
	ctx.orderService.ClearPromotions()
	return nil
}

func (ctx *OrderTestContext) theThresholdDiscountPromotionIsConfigured(table *godog.Table) error {
	for _, row := range table.Rows[1:] { // 跳過標題行
		threshold, err := strconv.Atoi(row.Cells[0].Value)
		if err != nil {
			return fmt.Errorf("invalid threshold: %s", row.Cells[0].Value)
		}
		discount, err := strconv.Atoi(row.Cells[1].Value)
		if err != nil {
			return fmt.Errorf("invalid discount: %s", row.Cells[1].Value)
		}

		promotion := &domain.ThresholdDiscountPromotion{
			Threshold: threshold,
			Discount:  discount,
		}
		ctx.orderService.AddPromotion(promotion)
	}
	return nil
}

func (ctx *OrderTestContext) theBuyOneGetOnePromotionForCosmeticsIsActive() error {
	promotion := &domain.BuyOneGetOnePromotion{
		Category: "cosmetics",
	}
	ctx.orderService.AddPromotion(promotion)
	return nil
}

func (ctx *OrderTestContext) aCustomerPlacesAnOrderWith(table *godog.Table) error {
	ctx.orderItems = make([]domain.OrderItem, 0)

	// 取得標題行以判斷欄位位置
	headers := make(map[string]int)
	for i, cell := range table.Rows[0].Cells {
		headers[cell.Value] = i
	}

	for _, row := range table.Rows[1:] { // 跳過標題行
		productName := row.Cells[headers["productName"]].Value

		var quantity int
		var err error
		if quantityIndex, exists := headers["quantity"]; exists {
			quantity, err = strconv.Atoi(row.Cells[quantityIndex].Value)
			if err != nil {
				return fmt.Errorf("invalid quantity: %s", row.Cells[quantityIndex].Value)
			}
		}

		var unitPrice int
		if unitPriceIndex, exists := headers["unitPrice"]; exists {
			unitPrice, err = strconv.Atoi(row.Cells[unitPriceIndex].Value)
			if err != nil {
				return fmt.Errorf("invalid unit price: %s", row.Cells[unitPriceIndex].Value)
			}
		}

		// 根據商品名稱判斷類別
		category := "apparel" // 預設類別
		if categoryIndex, exists := headers["category"]; exists {
			category = row.Cells[categoryIndex].Value
		} else {
			// 根據商品名稱推斷類別
			if productName == "口紅" || productName == "粉底液" {
				category = "cosmetics"
			}
		}

		product := domain.Product{
			Name:     productName,
			Category: category,
			Price:    unitPrice,
		}

		orderItem := domain.OrderItem{
			Product:  product,
			Quantity: quantity,
		}

		ctx.orderItems = append(ctx.orderItems, orderItem)
	}

	// 處理訂單
	ctx.order, ctx.lastError = ctx.orderService.ProcessOrder(ctx.orderItems)
	return ctx.lastError
}

func (ctx *OrderTestContext) theOrderSummaryShouldBe(table *godog.Table) error {
	if ctx.order == nil {
		return fmt.Errorf("no order processed")
	}

	summary := ctx.orderService.GetOrderSummary(ctx.order)

	for _, row := range table.Rows[1:] { // 跳過標題行
		for i, cell := range row.Cells {
			header := table.Rows[0].Cells[i].Value
			expectedValue, err := strconv.Atoi(cell.Value)
			if err != nil {
				return fmt.Errorf("invalid expected value for %s: %s", header, cell.Value)
			}

			var actualValue int
			switch header {
			case "totalAmount":
				actualValue = summary.TotalAmount
			case "originalAmount":
				actualValue = summary.OriginalAmount
			case "discount":
				actualValue = summary.Discount
			default:
				return fmt.Errorf("unknown header: %s", header)
			}

			if actualValue != expectedValue {
				return fmt.Errorf("expected %s to be %d, but got %d", header, expectedValue, actualValue)
			}
		}
	}

	return nil
}

func (ctx *OrderTestContext) theCustomerShouldReceive(table *godog.Table) error {
	if ctx.order == nil {
		return fmt.Errorf("no order processed")
	}

	receivedItems := ctx.orderService.GetReceivedItems(ctx.order)

	if len(receivedItems) != len(table.Rows)-1 {
		return fmt.Errorf("expected %d items, but got %d", len(table.Rows)-1, len(receivedItems))
	}

	for i, row := range table.Rows[1:] { // 跳過標題行
		expectedProductName := row.Cells[0].Value
		expectedQuantity, err := strconv.Atoi(row.Cells[1].Value)
		if err != nil {
			return fmt.Errorf("invalid expected quantity: %s", row.Cells[1].Value)
		}

		if i >= len(receivedItems) {
			return fmt.Errorf("missing received item for %s", expectedProductName)
		}

		receivedItem := receivedItems[i]
		if receivedItem.ProductName != expectedProductName {
			return fmt.Errorf("expected product %s, but got %s", expectedProductName, receivedItem.ProductName)
		}

		if receivedItem.Quantity != expectedQuantity {
			return fmt.Errorf("expected quantity %d for %s, but got %d", expectedQuantity, expectedProductName, receivedItem.Quantity)
		}
	}

	return nil
}
