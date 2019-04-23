package types

import "math/big"

const (
	AergoSystem     = "aergo.system"
	AergoName       = "aergo.name"
	AergoEnterprise = "aergo.enterprise"
)

const MaxCandidates = 30

const CreateProposal = "v1createProposal"
const VoteProposal = "v1voteProposal"
const VoteBP = "v1voteBP"

const proposalPrefixKey = "proposal"

var AllVotes = [...]string{VoteBP}

func (vl VoteList) Len() int { return len(vl.Votes) }
func (vl VoteList) Less(i, j int) bool {
	result := new(big.Int).SetBytes(vl.Votes[i].Amount).Cmp(new(big.Int).SetBytes(vl.Votes[j].Amount))
	if result == -1 {
		return true
	} else if result == 0 {
		if len(vl.Votes[i].Candidate) == 39 /*peer id length*/ {
			return new(big.Int).SetBytes(vl.Votes[i].Candidate[7:]).Cmp(new(big.Int).SetBytes(vl.Votes[j].Candidate[7:])) > 0
		}
		return new(big.Int).SetBytes(vl.Votes[i].Candidate).Cmp(new(big.Int).SetBytes(vl.Votes[j].Candidate)) > 0
	}
	return false
}
func (vl VoteList) Swap(i, j int) { vl.Votes[i], vl.Votes[j] = vl.Votes[j], vl.Votes[i] }

func (v *Vote) GetAmountBigInt() *big.Int {
	return new(big.Int).SetBytes(v.Amount)
}

func (a *Proposal) GetKey() []byte {
	return []byte(agendaPrefixKey + "\\" + a.Id)
}

func GenAgendaKey(name, version string) []byte {
	return []byte(agendaPrefixKey + "\\" + name)
}
