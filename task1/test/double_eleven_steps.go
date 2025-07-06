package test

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cucumber/godog"
	"github.com/mattli001/AI-100x-SE-Join-Quest/task1/src/domain"
	"github.com/mattli001/AI-100x-SE-Join-Quest/task1/src/service"
)

// DoubleElevenTestContext 雙十一測試上下文
type DoubleElevenTestContext struct {
	orderService *service.OrderService
	orderItems   []domain.OrderItem
	order        *domain.Order
	lastError    error
}

// RegisterDoubleElevenSteps 註冊雙十一測試步驟
func (ctx *DoubleElevenTestContext) RegisterDoubleElevenSteps(sc *godog.ScenarioContext) {
	sc.Before(ctx.beforeScenario)

	// Given steps
	sc.Step(`^the Double Eleven promotion is active$`, ctx.theDoubleElevenPromotionIsActive)

	// When steps
	sc.Step(`^a customer places an order with:$`, ctx.aCustomerPlacesAnOrderWith)

	// Then steps
	sc.Step(`^the order summary should be:$`, ctx.theOrderSummaryShouldBe)
	sc.Step(`^the discount details should be:$`, ctx.theDiscountDetailsShouldBe)
}

func (ctx *DoubleElevenTestContext) beforeScenario(ctx2 context.Context, sc *godog.Scenario) (context.Context, error) {
	ctx.orderService = service.NewOrderService()
	ctx.orderItems = make([]domain.OrderItem, 0)
	ctx.order = nil
	ctx.lastError = nil
	return ctx2, nil
}

func (ctx *DoubleElevenTestContext) theDoubleElevenPromotionIsActive() error {
	// 建立雙十一促銷
	promotion := domain.NewDoubleElevenPromotion()
	ctx.orderService.AddPromotion(promotion)
	return nil
}

func (ctx *DoubleElevenTestContext) aCustomerPlacesAnOrderWith(table *godog.Table) error {
	ctx.orderItems = make([]domain.OrderItem, 0)

	// 取得標題行以判斷欄位位置
	headers := make(map[string]int)
	for i, cell := range table.Rows[0].Cells {
		headers[cell.Value] = i
	}

	for _, row := range table.Rows[1:] { // 跳過標題行
		productName := row.Cells[headers["productName"]].Value

		quantity, err := strconv.Atoi(row.Cells[headers["quantity"]].Value)
		if err != nil {
			return fmt.Errorf("invalid quantity: %s", row.Cells[headers["quantity"]].Value)
		}

		unitPrice, err := strconv.Atoi(row.Cells[headers["unitPrice"]].Value)
		if err != nil {
			return fmt.Errorf("invalid unit price: %s", row.Cells[headers["unitPrice"]].Value)
		}

		product := domain.Product{
			Name:     productName,
			Category: "apparel",
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

func (ctx *DoubleElevenTestContext) theOrderSummaryShouldBe(table *godog.Table) error {
	if ctx.order == nil {
		return fmt.Errorf("no order processed")
	}

	summary := ctx.orderService.GetOrderSummary(ctx.order)

	for _, row := range table.Rows[1:] { // 跳過標題行
		expectedTotalAmount, err := strconv.Atoi(row.Cells[0].Value)
		if err != nil {
			return fmt.Errorf("invalid expected total amount: %s", row.Cells[0].Value)
		}

		if summary.TotalAmount != expectedTotalAmount {
			return fmt.Errorf("expected total amount to be %d, but got %d", expectedTotalAmount, summary.TotalAmount)
		}
	}

	return nil
}

func (ctx *DoubleElevenTestContext) theDiscountDetailsShouldBe(table *godog.Table) error {
	if ctx.order == nil {
		return fmt.Errorf("no order processed")
	}

	// 取得折扣詳情但不實作具體邏輯
	discountDetails := ctx.orderService.GetDiscountDetails(ctx.order)

	// 取得標題行以判斷欄位位置
	headers := make(map[string]int)
	for i, cell := range table.Rows[0].Cells {
		headers[cell.Value] = i
	}

	for i, row := range table.Rows[1:] { // 跳過標題行
		if i >= len(discountDetails) {
			return fmt.Errorf("missing discount detail for row %d", i+1)
		}

		detail := discountDetails[i]
		expectedProductName := row.Cells[headers["productName"]].Value

		if detail.ProductName != expectedProductName {
			return fmt.Errorf("expected product name %s, but got %s", expectedProductName, detail.ProductName)
		}

		expectedDiscountedQuantity, _ := strconv.Atoi(row.Cells[headers["discountedQuantity"]].Value)
		if detail.DiscountedQuantity != expectedDiscountedQuantity {
			return fmt.Errorf("expected discounted quantity %d, but got %d", expectedDiscountedQuantity, detail.DiscountedQuantity)
		}

		expectedRegularQuantity, _ := strconv.Atoi(row.Cells[headers["regularQuantity"]].Value)
		if detail.RegularQuantity != expectedRegularQuantity {
			return fmt.Errorf("expected regular quantity %d, but got %d", expectedRegularQuantity, detail.RegularQuantity)
		}

		expectedDiscountedAmount, _ := strconv.Atoi(row.Cells[headers["discountedAmount"]].Value)
		if detail.DiscountedAmount != expectedDiscountedAmount {
			return fmt.Errorf("expected discounted amount %d, but got %d", expectedDiscountedAmount, detail.DiscountedAmount)
		}

		expectedRegularAmount, _ := strconv.Atoi(row.Cells[headers["regularAmount"]].Value)
		if detail.RegularAmount != expectedRegularAmount {
			return fmt.Errorf("expected regular amount %d, but got %d", expectedRegularAmount, detail.RegularAmount)
		}
	}

	return nil
}
