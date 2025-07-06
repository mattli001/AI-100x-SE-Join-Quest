package chinesechess

type Position struct {
	Row int
	Col int
}

type Move struct {
	From Position
	To   Position
}

type Piece struct {
	Color string // "Red" or "Black"
	Type  string // "General", "Guard", "Rook", "Horse", "Cannon", "Elephant", "Soldier"
}

type Board struct {
	pieces             map[Position]*Piece
	lastCapturedGeneral *Piece
}

func NewBoard() *Board {
	return &Board{
		pieces:             make(map[Position]*Piece),
		lastCapturedGeneral: nil,
	}
}

func (b *Board) PlacePiece(row, col int, color, pieceType string) {
	pos := Position{Row: row, Col: col}
	piece := &Piece{
		Color: color,
		Type:  pieceType,
	}
	b.pieces[pos] = piece
}

func (b *Board) GetPiece(pos Position) *Piece {
	return b.pieces[pos]
}

func (b *Board) MakeMove(move Move) *MoveResult {
	// 取得起始位置的棋子
	piece := b.GetPiece(move.From)
	if piece == nil {
		return &MoveResult{
			IsLegal: false,
			Error:   "no piece at source position",
		}
	}
	
	// 根據棋子類型進行移動驗證
	var result *MoveResult
	switch piece.Type {
	case "General":
		result = b.validateGeneralMove(move, piece)
	case "Guard":
		result = b.validateGuardMove(move, piece)
	case "Rook":
		result = b.validateRookMove(move, piece)
	case "Horse":
		result = b.validateHorseMove(move, piece)
	case "Cannon":
		result = b.validateCannonMove(move, piece)
	case "Elephant":
		result = b.validateElephantMove(move, piece)
	case "Soldier":
		result = b.validateSoldierMove(move, piece)
	default:
		return &MoveResult{
			IsLegal: false,
			Error:   "move validation not implemented for this piece type",
		}
	}
	
	// 如果移動合法，實際執行移動
	if result.IsLegal {
		// 檢查目標位置是否有敵方棋子被捕獲
		capturedPiece := b.GetPiece(move.To)
		
		// 移除起始位置的棋子
		delete(b.pieces, move.From)
		// 將棋子放到目標位置（如果目標位置有敵方棋子會被覆蓋，即捕獲）
		b.pieces[move.To] = piece
		
		// 檢查是否捕獲了將/帥，如果是則記錄勝負結果
		if capturedPiece != nil && capturedPiece.Type == "General" {
			b.lastCapturedGeneral = capturedPiece
		} else {
			b.lastCapturedGeneral = nil
		}
	}
	
	return result
}

func (b *Board) validateGeneralMove(move Move, piece *Piece) *MoveResult {
	// 計算移動距離
	rowDiff := abs(move.To.Row - move.From.Row)
	colDiff := abs(move.To.Col - move.From.Col)
	
	// 將/帥只能橫向或縱向移動一格
	if (rowDiff == 1 && colDiff == 0) || (rowDiff == 0 && colDiff == 1) {
		// 檢查是否在九宮格內
		if b.isInPalace(move.To, piece.Color) {
			// 檢查是否與對方將/帥照面
			if !b.wouldGeneralsFaceEachOther(move) {
				return &MoveResult{
					IsLegal: true,
					Error:   "",
				}
			}
			return &MoveResult{
				IsLegal: false,
				Error:   "generals cannot face each other",
			}
		}
		return &MoveResult{
			IsLegal: false,
			Error:   "general must stay within palace",
		}
	}
	
	return &MoveResult{
		IsLegal: false,
		Error:   "general can only move one step horizontally or vertically",
	}
}

func (b *Board) validateGuardMove(move Move, piece *Piece) *MoveResult {
	// 計算移動距離
	rowDiff := abs(move.To.Row - move.From.Row)
	colDiff := abs(move.To.Col - move.From.Col)
	
	// 士/仕只能斜向移動一格
	if rowDiff == 1 && colDiff == 1 {
		// 檢查是否在九宮格內
		if b.isInPalace(move.To, piece.Color) {
			return &MoveResult{
				IsLegal: true,
				Error:   "",
			}
		}
		return &MoveResult{
			IsLegal: false,
			Error:   "guard must stay within palace",
		}
	}
	
	return &MoveResult{
		IsLegal: false,
		Error:   "guard can only move one step diagonally",
	}
}

func (b *Board) validateRookMove(move Move, piece *Piece) *MoveResult {
	// 車只能沿著橫線或豎線移動
	if move.From.Row != move.To.Row && move.From.Col != move.To.Col {
		return &MoveResult{
			IsLegal: false,
			Error:   "rook can only move horizontally or vertically",
		}
	}
	
	// 檢查路徑上是否有阻擋
	if b.isPathBlocked(move.From, move.To) {
		return &MoveResult{
			IsLegal: false,
			Error:   "path is blocked",
		}
	}
	
	// 檢查目標位置是否有己方棋子
	targetPiece := b.GetPiece(move.To)
	if targetPiece != nil && targetPiece.Color == piece.Color {
		return &MoveResult{
			IsLegal: false,
			Error:   "cannot capture own piece",
		}
	}
	
	return &MoveResult{
		IsLegal: true,
		Error:   "",
	}
}

func (b *Board) validateHorseMove(move Move, piece *Piece) *MoveResult {
	rowDiff := abs(move.To.Row - move.From.Row)
	colDiff := abs(move.To.Col - move.From.Col)
	
	// 馬走日字：(2,1) 或 (1,2)
	if !((rowDiff == 2 && colDiff == 1) || (rowDiff == 1 && colDiff == 2)) {
		return &MoveResult{
			IsLegal: false,
			Error:   "horse must move in L-shape",
		}
	}
	
	// 檢查馬腳位置是否被阻擋
	var legPos Position
	if rowDiff == 2 {
		// 垂直移動2格，馬腳在中間的垂直位置
		legRow := move.From.Row
		if move.To.Row > move.From.Row {
			legRow = move.From.Row + 1
		} else {
			legRow = move.From.Row - 1
		}
		legPos = Position{Row: legRow, Col: move.From.Col}
	} else {
		// 水平移動2格，馬腳在中間的水平位置
		legCol := move.From.Col
		if move.To.Col > move.From.Col {
			legCol = move.From.Col + 1
		} else {
			legCol = move.From.Col - 1
		}
		legPos = Position{Row: move.From.Row, Col: legCol}
	}
	
	// 檢查馬腳位置是否有棋子（蹩腳）
	if b.GetPiece(legPos) != nil {
		return &MoveResult{
			IsLegal: false,
			Error:   "horse is blocked by adjacent piece",
		}
	}
	
	// 檢查目標位置是否有己方棋子
	targetPiece := b.GetPiece(move.To)
	if targetPiece != nil && targetPiece.Color == piece.Color {
		return &MoveResult{
			IsLegal: false,
			Error:   "cannot capture own piece",
		}
	}
	
	return &MoveResult{
		IsLegal: true,
		Error:   "",
	}
}

func (b *Board) isPathBlocked(from, to Position) bool {
	// 計算移動方向
	rowStep := 0
	colStep := 0
	
	if to.Row > from.Row {
		rowStep = 1
	} else if to.Row < from.Row {
		rowStep = -1
	}
	
	if to.Col > from.Col {
		colStep = 1
	} else if to.Col < from.Col {
		colStep = -1
	}
	
	// 檢查路徑上的每一格（不包括起點和終點）
	currentRow := from.Row + rowStep
	currentCol := from.Col + colStep
	
	for currentRow != to.Row || currentCol != to.Col {
		pos := Position{Row: currentRow, Col: currentCol}
		if b.GetPiece(pos) != nil {
			return true // 路徑被阻擋
		}
		currentRow += rowStep
		currentCol += colStep
	}
	
	return false // 路徑暢通
}

func (b *Board) isInPalace(pos Position, color string) bool {
	// 紅方九宮格: row 1-3, col 4-6
	// 黑方九宮格: row 8-10, col 4-6
	if color == "Red" {
		return pos.Row >= 1 && pos.Row <= 3 && pos.Col >= 4 && pos.Col <= 6
	} else if color == "Black" {
		return pos.Row >= 8 && pos.Row <= 10 && pos.Col >= 4 && pos.Col <= 6
	}
	return false
}

func (b *Board) wouldGeneralsFaceEachOther(move Move) bool {
	// 建立移動後的臨時棋盤狀態
	tempBoard := make(map[Position]*Piece)
	for pos, piece := range b.pieces {
		tempBoard[pos] = piece
	}
	
	// 執行移動
	piece := tempBoard[move.From]
	delete(tempBoard, move.From)
	tempBoard[move.To] = piece
	
	// 尋找兩個將/帥的位置
	var redGeneral, blackGeneral *Position
	for pos, p := range tempBoard {
		if p.Type == "General" {
			if p.Color == "Red" {
				redGeneral = &pos
			} else if p.Color == "Black" {
				blackGeneral = &pos
			}
		}
	}
	
	// 如果沒有找到兩個將/帥，則不會照面
	if redGeneral == nil || blackGeneral == nil {
		return false
	}
	
	// 檢查是否在同一列
	if redGeneral.Col == blackGeneral.Col {
		// 檢查兩個將/帥之間是否有其他棋子
		minRow := min(redGeneral.Row, blackGeneral.Row)
		maxRow := max(redGeneral.Row, blackGeneral.Row)
		
		for row := minRow + 1; row < maxRow; row++ {
			pos := Position{Row: row, Col: redGeneral.Col}
			if tempBoard[pos] != nil {
				return false // 有棋子阻擋，不會照面
			}
		}
		return true // 沒有阻擋，會照面
	}
	
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (b *Board) validateElephantMove(move Move, piece *Piece) *MoveResult {
	rowDiff := abs(move.To.Row - move.From.Row)
	colDiff := abs(move.To.Col - move.From.Col)
	
	// 象走田字：必須斜向移動2格
	if rowDiff != 2 || colDiff != 2 {
		return &MoveResult{
			IsLegal: false,
			Error:   "elephant must move 2 steps diagonally",
		}
	}
	
	// 檢查是否過河
	if b.wouldElephantCrossRiver(move.To, piece.Color) {
		return &MoveResult{
			IsLegal: false,
			Error:   "elephant cannot cross the river",
		}
	}
	
	// 檢查象眼位置是否被阻擋
	midRow := (move.From.Row + move.To.Row) / 2
	midCol := (move.From.Col + move.To.Col) / 2
	midPos := Position{Row: midRow, Col: midCol}
	
	if b.GetPiece(midPos) != nil {
		return &MoveResult{
			IsLegal: false,
			Error:   "elephant is blocked by piece at midpoint",
		}
	}
	
	// 檢查目標位置是否有己方棋子
	targetPiece := b.GetPiece(move.To)
	if targetPiece != nil && targetPiece.Color == piece.Color {
		return &MoveResult{
			IsLegal: false,
			Error:   "cannot capture own piece",
		}
	}
	
	return &MoveResult{
		IsLegal: true,
		Error:   "",
	}
}

func (b *Board) wouldElephantCrossRiver(pos Position, color string) bool {
	// 中國象棋中，河界在第5行和第6行之間
	// 紅方象不能越過第5行，黑方象不能越過第6行
	if color == "Red" {
		return pos.Row > 5
	} else if color == "Black" {
		return pos.Row < 6
	}
	return false
}

func (b *Board) validateSoldierMove(move Move, piece *Piece) *MoveResult {
	rowDiff := move.To.Row - move.From.Row
	colDiff := move.To.Col - move.From.Col
	
	// 計算移動距離
	absRowDiff := abs(rowDiff)
	absColDiff := abs(colDiff)
	
	// 兵/卒只能移動一格
	if absRowDiff > 1 || absColDiff > 1 || (absRowDiff + absColDiff) != 1 {
		return &MoveResult{
			IsLegal: false,
			Error:   "soldier can only move one step",
		}
	}
	
	// 檢查是否已經過河
	hasCrossedRiver := b.hasSoldierCrossedRiver(move.From, piece.Color)
	
	if !hasCrossedRiver {
		// 未過河：只能向前移動
		if piece.Color == "Red" {
			// 紅方兵只能向上（row增加）
			if rowDiff != 1 || colDiff != 0 {
				return &MoveResult{
					IsLegal: false,
					Error:   "soldier can only move forward before crossing river",
				}
			}
		} else if piece.Color == "Black" {
			// 黑方卒只能向下（row減少）
			if rowDiff != -1 || colDiff != 0 {
				return &MoveResult{
					IsLegal: false,
					Error:   "soldier can only move forward before crossing river",
				}
			}
		}
	} else {
		// 已過河：可以向前或左右移動，但不能向後
		if piece.Color == "Red" {
			// 紅方兵不能向下（row減少）
			if rowDiff < 0 {
				return &MoveResult{
					IsLegal: false,
					Error:   "soldier cannot move backward after crossing river",
				}
			}
		} else if piece.Color == "Black" {
			// 黑方卒不能向上（row增加）
			if rowDiff > 0 {
				return &MoveResult{
					IsLegal: false,
					Error:   "soldier cannot move backward after crossing river",
				}
			}
		}
	}
	
	// 檢查目標位置是否有己方棋子
	targetPiece := b.GetPiece(move.To)
	if targetPiece != nil && targetPiece.Color == piece.Color {
		return &MoveResult{
			IsLegal: false,
			Error:   "cannot capture own piece",
		}
	}
	
	return &MoveResult{
		IsLegal: true,
		Error:   "",
	}
}

func (b *Board) hasSoldierCrossedRiver(pos Position, color string) bool {
	// 紅方兵過河是指行數 > 5，黑方卒過河是指行數 < 6
	if color == "Red" {
		return pos.Row > 5
	} else if color == "Black" {
		return pos.Row < 6
	}
	return false
}

func (b *Board) validateCannonMove(move Move, piece *Piece) *MoveResult {
	// 炮只能沿著橫線或豎線移動
	if move.From.Row != move.To.Row && move.From.Col != move.To.Col {
		return &MoveResult{
			IsLegal: false,
			Error:   "cannon can only move horizontally or vertically",
		}
	}
	
	// 檢查目標位置是否有棋子
	targetPiece := b.GetPiece(move.To)
	
	if targetPiece == nil {
		// 目標位置沒有棋子，炮移動如同車（路徑必須暢通）
		if b.isPathBlocked(move.From, move.To) {
			return &MoveResult{
				IsLegal: false,
				Error:   "path is blocked",
			}
		}
		return &MoveResult{
			IsLegal: true,
			Error:   "",
		}
	} else {
		// 目標位置有棋子，檢查是否為敵方棋子
		if targetPiece.Color == piece.Color {
			return &MoveResult{
				IsLegal: false,
				Error:   "cannot capture own piece",
			}
		}
		
		// 炮吃子必須跳過恰好一個棋子（炮台）
		screenCount := b.countPiecesInPath(move.From, move.To)
		if screenCount != 1 {
			if screenCount == 0 {
				return &MoveResult{
					IsLegal: false,
					Error:   "cannon must jump over exactly one piece to capture",
				}
			} else {
				return &MoveResult{
					IsLegal: false,
					Error:   "cannon can only jump over one piece",
				}
			}
		}
		
		return &MoveResult{
			IsLegal: true,
			Error:   "",
		}
	}
}

func (b *Board) countPiecesInPath(from, to Position) int {
	count := 0
	
	// 計算移動方向
	rowStep := 0
	colStep := 0
	
	if to.Row > from.Row {
		rowStep = 1
	} else if to.Row < from.Row {
		rowStep = -1
	}
	
	if to.Col > from.Col {
		colStep = 1
	} else if to.Col < from.Col {
		colStep = -1
	}
	
	// 計算路徑上的棋子數量（不包括起點和終點）
	currentRow := from.Row + rowStep
	currentCol := from.Col + colStep
	
	for currentRow != to.Row || currentCol != to.Col {
		pos := Position{Row: currentRow, Col: currentCol}
		if b.GetPiece(pos) != nil {
			count++
		}
		currentRow += rowStep
		currentCol += colStep
	}
	
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (b *Board) CheckGameResult() *GameResult {
	// 檢查是否在上一次移動中捕獲了將/帥
	if b.lastCapturedGeneral != nil {
		// 捕獲了將/帥，遊戲結束
		if b.lastCapturedGeneral.Color == "Red" {
			return &GameResult{
				IsGameOver: true,
				Winner:     "Black",
			}
		} else if b.lastCapturedGeneral.Color == "Black" {
			return &GameResult{
				IsGameOver: true,
				Winner:     "Red",
			}
		}
	}
	
	// 沒有捕獲將/帥，遊戲繼續
	return &GameResult{
		IsGameOver: false,
		Winner:     "",
	}
}

type MoveResult struct {
	IsLegal bool
	Error   string
}

type GameResult struct {
	IsGameOver bool
	Winner     string
}