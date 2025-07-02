@double_eleven_promotion
Feature: Double Eleven Promotion
  As a shopper
  I want to enjoy Double Eleven promotion discounts
  So that I can get discounts when buying 10 or more of the same product

  Background:
    Given the Double Eleven promotion is active

  Scenario: Same product with 12 items - partial discount
    When a customer places an order with:
      | productName | quantity | unitPrice |
      | 襪子          | 12       | 100       |
    Then the order summary should be:
      | totalAmount |
      | 1000        |
    And the discount details should be:
      | productName | discountedQuantity | regularQuantity | discountedAmount | regularAmount |
      | 襪子          | 10                 | 2               | 800              | 200           |

  Scenario: Same product with 27 items - multiple discount groups
    When a customer places an order with:
      | productName | quantity | unitPrice |
      | 襪子          | 27       | 100       |
    Then the order summary should be:
      | totalAmount |
      | 2300        |
    And the discount details should be:
      | productName | discountedQuantity | regularQuantity | discountedAmount | regularAmount |
      | 襪子          | 20                 | 7               | 1600             | 700           |

  Scenario: Ten different products - no discount applies
    When a customer places an order with:
      | productName | quantity | unitPrice |
      | 商品A         | 1        | 100       |
      | 商品B         | 1        | 100       |
      | 商品C         | 1        | 100       |
      | 商品D         | 1        | 100       |
      | 商品E         | 1        | 100       |
      | 商品F         | 1        | 100       |
      | 商品G         | 1        | 100       |
      | 商品H         | 1        | 100       |
      | 商品I         | 1        | 100       |
      | 商品J         | 1        | 100       |
    Then the order summary should be:
      | totalAmount |
      | 1000        |
    And the discount details should be:
      | productName | discountedQuantity | regularQuantity | discountedAmount | regularAmount |
      | 商品A         | 0                  | 1               | 0                | 100           |
      | 商品B         | 0                  | 1               | 0                | 100           |
      | 商品C         | 0                  | 1               | 0                | 100           |
      | 商品D         | 0                  | 1               | 0                | 100           |
      | 商品E         | 0                  | 1               | 0                | 100           |
      | 商品F         | 0                  | 1               | 0                | 100           |
      | 商品G         | 0                  | 1               | 0                | 100           |
      | 商品H         | 0                  | 1               | 0                | 100           |
      | 商品I         | 0                  | 1               | 0                | 100           |
      | 商品J         | 0                  | 1               | 0                | 100           | 