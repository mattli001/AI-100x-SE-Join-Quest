package chinesechess

// PieceType 棋子類型
type PieceType int

const (
	General PieceType = iota // 將/帥
	Guard                    // 士/仕
	Elephant                 // 相/象
	Horse                    // 馬/傌
	Rook                     // 車
	Cannon                   // 炮/砲
	Soldier                  // 兵/卒
)

// Color 棋子顏色
type Color int

const (
	Red   Color = iota // 紅方
	Black              // 黑方
)

// Piece 棋子
type Piece struct {
	Type  PieceType
	Color Color
}

// Position 位置
type Position struct {
	Row int
	Col int
}

// Board 棋盤
type Board struct {
	pieces map[Position]*Piece
}

// NewBoard 建立新棋盤
func NewBoard() *Board {
	return &Board{
		pieces: make(map[Position]*Piece),
	}
}

// SetPiece 在指定位置放置棋子
func (b *Board) SetPiece(pos Position, piece *Piece) {
	b.pieces[pos] = piece
}

// GetPiece 取得指定位置的棋子
func (b *Board) GetPiece(pos Position) *Piece {
	return b.pieces[pos]
}

// IsValidMove 檢查移動是否合法
func (b *Board) IsValidMove(from, to Position) bool {
	valid, _ := b.IsValidMoveWithResult(from, to)
	return valid
}

// IsValidMoveWithResult 檢查移動是否合法並回傳遊戲結果
func (b *Board) IsValidMoveWithResult(from, to Position) (bool, string) {
	piece := b.GetPiece(from)
	if piece == nil {
		return false, ""
	}
	
	// 檢查基本移動規則
	var isValid bool
	switch piece.Type {
	case General:
		isValid = b.isValidGeneralMove(from, to, piece.Color)
	case Guard:
		isValid = b.isValidGuardMove(from, to, piece.Color)
	case Rook:
		isValid = b.isValidRookMove(from, to)
	case Horse:
		isValid = b.isValidHorseMove(from, to)
	case Cannon:
		isValid = b.isValidCannonMove(from, to)
	case Elephant:
		isValid = b.isValidElephantMove(from, to, piece.Color)
	case Soldier:
		isValid = b.isValidSoldierMove(from, to, piece.Color)
	default:
		return false, ""
	}
	
	if !isValid {
		return false, ""
	}
	
	// 檢查是否吃掉對方棋子
	targetPiece := b.GetPiece(to)
	if targetPiece != nil {
		// 不能吃自己的棋子
		if targetPiece.Color == piece.Color {
			return false, ""
		}
		// 如果吃掉對方將軍，遊戲結束
		if targetPiece.Type == General {
			if piece.Color == Red {
				return true, "red_wins"
			} else {
				return true, "black_wins"
			}
		}
	}
	
	return true, "continue"
}

// isValidGeneralMove 檢查將的移動是否合法
func (b *Board) isValidGeneralMove(from, to Position, color Color) bool {
	// 檢查是否在九宮內
	if !b.isInPalace(to, color) {
		return false
	}
	
	// 檢查是否只移動一格，且為上下左右
	rowDiff := abs(to.Row - from.Row)
	colDiff := abs(to.Col - from.Col)
	
	// 只能上下左右移動一格
	if !((rowDiff == 1 && colDiff == 0) || (rowDiff == 0 && colDiff == 1)) {
		return false
	}
	
	// 檢查兩將是否會相對
	if b.wouldGeneralsFaceEachOther(from, to, color) {
		return false
	}
	
	return true
}

// isValidGuardMove 檢查士的移動是否合法
func (b *Board) isValidGuardMove(from, to Position, color Color) bool {
	// 檢查是否在九宮內
	if !b.isInPalace(to, color) {
		return false
	}
	
	// 檢查是否只斜向移動一格
	rowDiff := abs(to.Row - from.Row)
	colDiff := abs(to.Col - from.Col)
	
	// 只能斜向移動一格
	if rowDiff == 1 && colDiff == 1 {
		return true
	}
	
	return false
}

// isInPalace 檢查位置是否在九宮內
func (b *Board) isInPalace(pos Position, color Color) bool {
	if color == Red {
		// 紅方九宮：行 4-6，列 1-3
		return pos.Row >= 1 && pos.Row <= 3 && pos.Col >= 4 && pos.Col <= 6
	} else {
		// 黑方九宮：行 4-6，列 8-10
		return pos.Row >= 8 && pos.Row <= 10 && pos.Col >= 4 && pos.Col <= 6
	}
}

// wouldGeneralsFaceEachOther 檢查兩將是否會相對
func (b *Board) wouldGeneralsFaceEachOther(from, to Position, color Color) bool {
	// 建立一個臨時棋盤狀態來模擬移動
	tempBoard := make(map[Position]*Piece)
	for pos, piece := range b.pieces {
		tempBoard[pos] = piece
	}
	
	// 執行移動
	delete(tempBoard, from)
	tempBoard[to] = b.pieces[from]
	
	// 找到兩將的位置
	var redGeneralPos, blackGeneralPos Position
	var foundRed, foundBlack bool
	
	for pos, piece := range tempBoard {
		if piece.Type == General {
			if piece.Color == Red {
				redGeneralPos = pos
				foundRed = true
			} else {
				blackGeneralPos = pos
				foundBlack = true
			}
		}
	}
	
	// 如果找不到任一將，則不相對
	if !foundRed || !foundBlack {
		return false
	}
	
	// 檢查是否在同一列
	if redGeneralPos.Col != blackGeneralPos.Col {
		return false
	}
	
	// 檢查兩將之間是否有其他棋子
	minRow := min(redGeneralPos.Row, blackGeneralPos.Row)
	maxRow := max(redGeneralPos.Row, blackGeneralPos.Row)
	
	for row := minRow + 1; row < maxRow; row++ {
		pos := Position{Row: row, Col: redGeneralPos.Col}
		if tempBoard[pos] != nil {
			return false // 有棋子阻擋，不相對
		}
	}
	
	return true // 兩將相對
}

// min 取得較小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max 取得較大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// isValidRookMove 檢查車的移動是否合法
func (b *Board) isValidRookMove(from, to Position) bool {
	// 車只能直線移動
	if from.Row != to.Row && from.Col != to.Col {
		return false
	}
	
	// 檢查路徑是否有阻擋
	return b.isPathClear(from, to)
}

// isValidHorseMove 檢查馬的移動是否合法
func (b *Board) isValidHorseMove(from, to Position) bool {
	rowDiff := abs(to.Row - from.Row)
	colDiff := abs(to.Col - from.Col)
	
	// 馬走日字
	if !((rowDiff == 2 && colDiff == 1) || (rowDiff == 1 && colDiff == 2)) {
		return false
	}
	
	// 檢查馬腿是否被阻擋
	var blockPos Position
	if rowDiff == 2 {
		// 縱向移動2格，檢查中間位置
		blockPos = Position{Row: from.Row + (to.Row-from.Row)/2, Col: from.Col}
	} else {
		// 橫向移動2格，檢查中間位置
		blockPos = Position{Row: from.Row, Col: from.Col + (to.Col-from.Col)/2}
	}
	
	return b.GetPiece(blockPos) == nil
}

// isValidCannonMove 檢查炮的移動是否合法
func (b *Board) isValidCannonMove(from, to Position) bool {
	// 炮只能直線移動
	if from.Row != to.Row && from.Col != to.Col {
		return false
	}
	
	targetPiece := b.GetPiece(to)
	piecesInBetween := b.countPiecesInBetween(from, to)
	
	if targetPiece == nil {
		// 不吃子的移動，路徑必須無阻
		return piecesInBetween == 0
	} else {
		// 吃子的移動，路徑中必須恰好有一個棋子作為炮架
		return piecesInBetween == 1
	}
}

// isValidElephantMove 檢查相的移動是否合法
func (b *Board) isValidElephantMove(from, to Position, color Color) bool {
	// 相不能過河
	if color == Red && to.Row > 5 {
		return false
	}
	if color == Black && to.Row < 6 {
		return false
	}
	
	// 相走田字
	rowDiff := abs(to.Row - from.Row)
	colDiff := abs(to.Col - from.Col)
	if rowDiff != 2 || colDiff != 2 {
		return false
	}
	
	// 檢查象眼是否被阻擋
	elephantEye := Position{
		Row: from.Row + (to.Row-from.Row)/2,
		Col: from.Col + (to.Col-from.Col)/2,
	}
	
	return b.GetPiece(elephantEye) == nil
}

// isValidSoldierMove 檢查兵的移動是否合法
func (b *Board) isValidSoldierMove(from, to Position, color Color) bool {
	rowDiff := to.Row - from.Row
	colDiff := abs(to.Col - from.Col)
	
	if color == Red {
		// 紅兵只能向前或橫向移動
		if rowDiff < 0 {
			return false // 不能後退
		}
		
		if from.Row <= 5 {
			// 未過河，只能向前
			return rowDiff == 1 && colDiff == 0
		} else {
			// 已過河，可以向前或橫向移動
			return (rowDiff == 1 && colDiff == 0) || (rowDiff == 0 && colDiff == 1)
		}
	} else {
		// 黑卒只能向下或橫向移動
		if rowDiff > 0 {
			return false // 不能後退
		}
		
		if from.Row >= 6 {
			// 未過河，只能向下
			return rowDiff == -1 && colDiff == 0
		} else {
			// 已過河，可以向下或橫向移動
			return (rowDiff == -1 && colDiff == 0) || (rowDiff == 0 && colDiff == 1)
		}
	}
}

// isPathClear 檢查兩點之間的路徑是否暢通
func (b *Board) isPathClear(from, to Position) bool {
	return b.countPiecesInBetween(from, to) == 0
}

// countPiecesInBetween 計算兩點之間的棋子數量
func (b *Board) countPiecesInBetween(from, to Position) int {
	count := 0
	
	if from.Row == to.Row {
		// 橫向移動
		minCol := min(from.Col, to.Col)
		maxCol := max(from.Col, to.Col)
		for col := minCol + 1; col < maxCol; col++ {
			if b.GetPiece(Position{Row: from.Row, Col: col}) != nil {
				count++
			}
		}
	} else if from.Col == to.Col {
		// 縱向移動
		minRow := min(from.Row, to.Row)
		maxRow := max(from.Row, to.Row)
		for row := minRow + 1; row < maxRow; row++ {
			if b.GetPiece(Position{Row: row, Col: from.Col}) != nil {
				count++
			}
		}
	}
	
	return count
}

// abs 計算絕對值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}