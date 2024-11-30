package tests

import (
	"rubrumcreation.com/go-paper-scissor/models"
	"testing"
)

func TestRock(t *testing.T) {
	rock := models.Rock{}
	rockType := rock.GetType() // Access exported GetType method

	if rockType != models.ROCK {
		t.Errorf("Expected ROCK, got %v", rock.GetType())
	}
	/*	paper := models.Paper{}
		scissor := models.Scissor{}

		if rock.Beats(paper) {
			t.Errorf("%v, should not beat: %v", rock.GetType(), paper.GetType())
		}

		if !rock.Beats(scissor) {
			t.Errorf("%v, should beat: %v", rock.GetType(), scissor.GetType())
		}
	*/
}

func TestPaper(t *testing.T) {
	paper := models.Paper{}
	paperType := paper.GetType()

	if paperType != models.PAPER {
		t.Errorf("Expected PAPER, got %v", paperType)
	}

	/*	rock := models.Rock{}
		scissor := models.Scissor{}
		if paper.Beats(scissor) {
			t.Errorf("%v, should not beat: %v", paper.GetType(), scissor.GetType())
		}

		if !paper.Beats(rock) {
			t.Errorf("%v, should beat: %v", paper.GetType(), rock.GetType())
		}
	*/
}

func TestScissor(t *testing.T) {
	scissor := models.Scissor{}
	scissorType := scissor.GetType()
	if scissorType != models.SCISSOR {
		t.Errorf("Expected SCISSOR, got %v", scissorType)
	}
	/*	rock := models.Rock{}
		paper := models.Paper{}

		if scissor.Beats(rock) {
			t.Errorf("%v, should not beat: %v", scissor.GetType(), rock.GetType())
		}

		if !scissor.Beats(paper) {
			t.Errorf("%v, should beat: %v", scissor.GetType(), paper.GetType())
		}
	*/
}
