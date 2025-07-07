package test

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/mattli001/AI-100x-SE-Join-Quest/task2-2/src/chinesechess"
	"github.com/cucumber/godog"
)

type Steps struct {
	// 測試狀態
	board *chinesechess.Board
	moveIsValid bool
	errorMessage string
	gameResult string
}

// 第一個最簡單的 scenario 的步驟
func (s *Steps) 棋盤為空僅有一枚紅將於(row, col int) error {
	s.board = chinesechess.NewBoard()
	pos := chinesechess.Position{Row: row, Col: col}
	piece := &chinesechess.Piece{
		Type:  chinesechess.General,
		Color: chinesechess.Red,
	}
	s.board.SetPiece(pos, piece)
	return nil
}

func (s *Steps) 紅方將從移動至(fromRow, fromCol, toRow, toCol int) error {
	from := chinesechess.Position{Row: fromRow, Col: fromCol}
	to := chinesechess.Position{Row: toRow, Col: toCol}
	s.moveIsValid = s.board.IsValidMove(from, to)
	return nil
}

func (s *Steps) 此步合法() error {
	if !s.moveIsValid {
		return fmt.Errorf("期望移動合法，但實際上不合法")
	}
	return nil
}

func (s *Steps) 此步不合法() error {
	if s.moveIsValid {
		return fmt.Errorf("期望移動不合法，但實際上合法")
	}
	return nil
}

func (s *Steps) 棋盤包含(table *godog.Table) error {
	s.board = chinesechess.NewBoard()
	
	// 跳過標題列
	for i := 1; i < len(table.Rows); i++ {
		row := table.Rows[i]
		if len(row.Cells) < 2 {
			continue
		}
		
		pieceStr := row.Cells[0].Value
		coordStr := row.Cells[1].Value
		
		// 解析座標 (row, col)
		coordStr = strings.Trim(coordStr, "()")
		parts := strings.Split(coordStr, ", ")
		if len(parts) != 2 {
			return fmt.Errorf("無法解析座標: %s", coordStr)
		}
		
		rowNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("無法解析行號: %s", parts[0])
		}
		
		colNum, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("無法解析列號: %s", parts[1])
		}
		
		pos := chinesechess.Position{Row: rowNum, Col: colNum}
		piece := s.parsePiece(pieceStr)
		if piece == nil {
			return fmt.Errorf("無法解析棋子: %s", pieceStr)
		}
		
		s.board.SetPiece(pos, piece)
	}
	
	return nil
}

func (s *Steps) parsePiece(pieceStr string) *chinesechess.Piece {
	switch pieceStr {
	case "紅將":
		return &chinesechess.Piece{Type: chinesechess.General, Color: chinesechess.Red}
	case "黑將":
		return &chinesechess.Piece{Type: chinesechess.General, Color: chinesechess.Black}
	case "紅士":
		return &chinesechess.Piece{Type: chinesechess.Guard, Color: chinesechess.Red}
	case "黑士":
		return &chinesechess.Piece{Type: chinesechess.Guard, Color: chinesechess.Black}
	case "紅車":
		return &chinesechess.Piece{Type: chinesechess.Rook, Color: chinesechess.Red}
	case "黑車":
		return &chinesechess.Piece{Type: chinesechess.Rook, Color: chinesechess.Black}
	case "紅馬":
		return &chinesechess.Piece{Type: chinesechess.Horse, Color: chinesechess.Red}
	case "黑馬":
		return &chinesechess.Piece{Type: chinesechess.Horse, Color: chinesechess.Black}
	case "紅炮":
		return &chinesechess.Piece{Type: chinesechess.Cannon, Color: chinesechess.Red}
	case "黑炮":
		return &chinesechess.Piece{Type: chinesechess.Cannon, Color: chinesechess.Black}
	case "紅相":
		return &chinesechess.Piece{Type: chinesechess.Elephant, Color: chinesechess.Red}
	case "黑象":
		return &chinesechess.Piece{Type: chinesechess.Elephant, Color: chinesechess.Black}
	case "紅兵":
		return &chinesechess.Piece{Type: chinesechess.Soldier, Color: chinesechess.Red}
	case "黑卒":
		return &chinesechess.Piece{Type: chinesechess.Soldier, Color: chinesechess.Black}
	default:
		return nil
	}
}

func (s *Steps) 棋盤為空僅有一枚紅士於(row, col int) error {
	s.board = chinesechess.NewBoard()
	pos := chinesechess.Position{Row: row, Col: col}
	piece := &chinesechess.Piece{
		Type:  chinesechess.Guard,
		Color: chinesechess.Red,
	}
	s.board.SetPiece(pos, piece)
	return nil
}

func (s *Steps) 紅方士從移動至(fromRow, fromCol, toRow, toCol int) error {
	from := chinesechess.Position{Row: fromRow, Col: fromCol}
	to := chinesechess.Position{Row: toRow, Col: toCol}
	s.moveIsValid = s.board.IsValidMove(from, to)
	return nil
}

// 車的步驟
func (s *Steps) 棋盤為空僅有一枚紅車於(row, col int) error {
	s.board = chinesechess.NewBoard()
	pos := chinesechess.Position{Row: row, Col: col}
	piece := &chinesechess.Piece{
		Type:  chinesechess.Rook,
		Color: chinesechess.Red,
	}
	s.board.SetPiece(pos, piece)
	return nil
}

func (s *Steps) 紅方車從移動至(fromRow, fromCol, toRow, toCol int) error {
	from := chinesechess.Position{Row: fromRow, Col: fromCol}
	to := chinesechess.Position{Row: toRow, Col: toCol}
	s.moveIsValid, s.gameResult = s.board.IsValidMoveWithResult(from, to)
	return nil
}

// 馬的步驟
func (s *Steps) 棋盤為空僅有一枚紅馬於(row, col int) error {
	s.board = chinesechess.NewBoard()
	pos := chinesechess.Position{Row: row, Col: col}
	piece := &chinesechess.Piece{
		Type:  chinesechess.Horse,
		Color: chinesechess.Red,
	}
	s.board.SetPiece(pos, piece)
	return nil
}

func (s *Steps) 紅方馬從移動至(fromRow, fromCol, toRow, toCol int) error {
	from := chinesechess.Position{Row: fromRow, Col: fromCol}
	to := chinesechess.Position{Row: toRow, Col: toCol}
	s.moveIsValid = s.board.IsValidMove(from, to)
	return nil
}

// 炮的步驟
func (s *Steps) 棋盤為空僅有一枚紅炮於(row, col int) error {
	s.board = chinesechess.NewBoard()
	pos := chinesechess.Position{Row: row, Col: col}
	piece := &chinesechess.Piece{
		Type:  chinesechess.Cannon,
		Color: chinesechess.Red,
	}
	s.board.SetPiece(pos, piece)
	return nil
}

func (s *Steps) 紅方炮從移動至(fromRow, fromCol, toRow, toCol int) error {
	from := chinesechess.Position{Row: fromRow, Col: fromCol}
	to := chinesechess.Position{Row: toRow, Col: toCol}
	s.moveIsValid = s.board.IsValidMove(from, to)
	return nil
}

// 相的步驟
func (s *Steps) 棋盤為空僅有一枚紅相於(row, col int) error {
	s.board = chinesechess.NewBoard()
	pos := chinesechess.Position{Row: row, Col: col}
	piece := &chinesechess.Piece{
		Type:  chinesechess.Elephant,
		Color: chinesechess.Red,
	}
	s.board.SetPiece(pos, piece)
	return nil
}

func (s *Steps) 紅方相從移動至(fromRow, fromCol, toRow, toCol int) error {
	from := chinesechess.Position{Row: fromRow, Col: fromCol}
	to := chinesechess.Position{Row: toRow, Col: toCol}
	s.moveIsValid = s.board.IsValidMove(from, to)
	return nil
}

// 兵的步驟
func (s *Steps) 棋盤為空僅有一枚紅兵於(row, col int) error {
	s.board = chinesechess.NewBoard()
	pos := chinesechess.Position{Row: row, Col: col}
	piece := &chinesechess.Piece{
		Type:  chinesechess.Soldier,
		Color: chinesechess.Red,
	}
	s.board.SetPiece(pos, piece)
	return nil
}

func (s *Steps) 紅方兵從移動至(fromRow, fromCol, toRow, toCol int) error {
	from := chinesechess.Position{Row: fromRow, Col: fromCol}
	to := chinesechess.Position{Row: toRow, Col: toCol}
	s.moveIsValid = s.board.IsValidMove(from, to)
	return nil
}

// 遊戲結果步驟
func (s *Steps) 紅方立即獲勝() error {
	if s.gameResult != "red_wins" {
		return fmt.Errorf("期望紅方立即獲勝，但實際結果為: %s", s.gameResult)
	}
	return nil
}

func (s *Steps) 僅此吃子並不結束遊戲() error {
	if s.gameResult == "red_wins" || s.gameResult == "black_wins" {
		return fmt.Errorf("期望遊戲繼續，但遊戲結束了: %s", s.gameResult)
	}
	return nil
}