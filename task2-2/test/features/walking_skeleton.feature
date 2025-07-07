Feature: Walking Skeleton - 中國象棋第一個測試

@General
Scenario: 紅將在九宮內移動（合法）
  Given 棋盤為空，僅有一枚紅將於 (1, 5)
  When 紅方將從 (1, 5) 移動至 (1, 4)
  Then 此步合法

@General
Scenario: 紅將離開九宮（不合法）
  Given 棋盤為空，僅有一枚紅將於 (1, 6)
  When 紅方將從 (1, 6) 移動至 (1, 7)
  Then 此步不合法

@General
Scenario: 兩將同列相對（不合法）
  Given 棋盤包含：
    | 棋子 | 座標 |
    | 紅將 | (2, 4) |
    | 黑將 | (8, 5) |
  When 紅方將從 (2, 4) 移動至 (2, 5)
  Then 此步不合法

@Guard
Scenario: 紅士在九宮內斜向移動（合法）
  Given 棋盤為空，僅有一枚紅士於 (1, 4)
  When 紅方士從 (1, 4) 移動至 (2, 5)
  Then 此步合法