package raffle

import (
	"fmt"
	"math/rand"
	"time"
)

type Raffle struct {
	Participants []*Participant
	Results      *Results
}

// add participant to raffle
func (r *Raffle) AddParticipant(name string, mail string) {
	r.Participants = append(r.Participants, &Participant{Name: name, Mail: mail})
}

func (r *Raffle) Draw() {
	// copying the participant list for mapping
	whomList := make([]*Participant, len(r.Participants))
	copy(whomList, r.Participants)
	// pairs it with a random person other than herself, standing on one by one of all the participants
	for _, p := range r.Participants {
		for i, np := range whomList {
			if p.Mail == np.Mail {
				whomList = append(whomList[:i], whomList[i+1:]...)
			}
		}
		sw := rand.NewSource(time.Now().UnixNano())
		rw := rand.New(sw)
		randomWhomIndex := rw.Intn(len(whomList))
		r.Results.newRelation(p, whomList[randomWhomIndex])
		if len(whomList) == randomWhomIndex+1 {
			whomList = whomList[:randomWhomIndex]
		} else {
			whomList = append(whomList[:randomWhomIndex], whomList[randomWhomIndex+1:]...)
		}
		whomList = append(whomList, p)
	}
}

// return new raffle struct
func NewRaffle() *Raffle {
	return &Raffle{
		Results: &Results{},
	}
}

type Participant struct {
	Name string
	Mail string
}

type Relation struct {
	Who  *Participant
	Whom *Participant
}

type Results struct {
	Relations []*Relation
}

// Print result (for convenience)
func (r *Results) Print() {
	for i, r := range r.Relations {
		fmt.Printf("%v - from: %v (%v) -> to: %v (%v) \n", i+1, r.Who.Name, r.Who.Mail, r.Whom.Name, r.Whom.Mail)
	}
}

// create a new relation and add to relations of raffle
func (r *Results) newRelation(who *Participant, whom *Participant) {
	r.Relations = append(r.Relations, &Relation{Who: who, Whom: whom})
}
