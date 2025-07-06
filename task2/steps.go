package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cucumber/godog"
	"chinese_chess/src"
)

type ChineseChessSteps struct {
	board       *chinesechess.Board
	moveResult  *chinesechess.MoveResult
	gameResult  *chinesechess.GameResult
}

func (s *ChineseChessSteps) theBoardIsEmptyExceptForAPieceAt(pieceType string, row, col string) error {
	r, _ := strconv.Atoi(row)
	c, _ := strconv.Atoi(col)
	
	s.board = chinesechess.NewBoard()
	
	// 解析棋子類型和顏色
	parts := strings.Split(pieceType, " ")
	if len(parts) != 2 {
		return fmt.Errorf("invalid piece format: %s", pieceType)
	}
	
	color := parts[0]
	piece := parts[1]
	
	// 放置棋子
	s.board.PlacePiece(r, c, color, piece)
	
	return nil
}

func (s *ChineseChessSteps) theBoardHas(table *godog.Table) error {
	s.board = chinesechess.NewBoard()
	
	for _, row := range table.Rows[1:] { // 跳過標題行
		pieceType := row.Cells[0].Value
		position := row.Cells[1].Value
		
		// 解析位置 (row, col)
		position = strings.Trim(position, "()")
		coords := strings.Split(position, ", ")
		r, _ := strconv.Atoi(coords[0])
		c, _ := strconv.Atoi(coords[1])
		
		// 解析棋子類型和顏色
		parts := strings.Split(pieceType, " ")
		if len(parts) != 2 {
			return fmt.Errorf("invalid piece format: %s", pieceType)
		}
		
		color := parts[0]
		piece := parts[1]
		
		// 放置棋子
		s.board.PlacePiece(r, c, color, piece)
	}
	
	return nil
}

func (s *ChineseChessSteps) playerMovesThePieceFromTo(player, pieceType, fromRow, fromCol, toRow, toCol string) error {
	fRow, _ := strconv.Atoi(fromRow)
	fCol, _ := strconv.Atoi(fromCol)
	tRow, _ := strconv.Atoi(toRow)
	tCol, _ := strconv.Atoi(toCol)
	
	// 執行移動並記錄結果
	move := chinesechess.Move{
		From: chinesechess.Position{Row: fRow, Col: fCol},
		To:   chinesechess.Position{Row: tRow, Col: tCol},
	}
	
	s.moveResult = s.board.MakeMove(move)
	
	// 檢查是否贏得遊戲
	s.gameResult = s.board.CheckGameResult()
	
	return nil
}

func (s *ChineseChessSteps) theMoveIsLegal() error {
	if s.moveResult == nil {
		return fmt.Errorf("no move result available")
	}
	
	if !s.moveResult.IsLegal {
		return fmt.Errorf("expected move to be legal, but it was illegal: %s", s.moveResult.Error)
	}
	
	return nil
}

func (s *ChineseChessSteps) theMoveIsIllegal() error {
	if s.moveResult == nil {
		return fmt.Errorf("no move result available")
	}
	
	if s.moveResult.IsLegal {
		return fmt.Errorf("expected move to be illegal, but it was legal")
	}
	
	return nil
}

func (s *ChineseChessSteps) playerWinsImmediately(player string) error {
	if s.gameResult == nil {
		return fmt.Errorf("no game result available")
	}
	
	if !s.gameResult.IsGameOver {
		return fmt.Errorf("expected game to be over, but it continues")
	}
	
	if s.gameResult.Winner != player {
		return fmt.Errorf("expected %s to win, but winner is %s", player, s.gameResult.Winner)
	}
	
	return nil
}

func (s *ChineseChessSteps) theGameIsNotOverJustFromThatCapture() error {
	if s.gameResult == nil {
		return fmt.Errorf("no game result available")
	}
	
	if s.gameResult.IsGameOver {
		return fmt.Errorf("expected game to continue, but it ended")
	}
	
	return nil
}