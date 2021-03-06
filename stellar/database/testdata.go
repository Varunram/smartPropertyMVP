package database

import (
	"fmt"
	"log"

	utils "github.com/YaleOpenLab/smartPropertyMVP/stellar/utils"
)

func InsertDummyData() error {
	var err error
	// populate database with dumym data
	var order1 Order

	var rec Recipient
	allRecs, err := RetrieveAllRecipients()
	if err != nil {
		log.Fatal(err)
	}
	if len(allRecs) == 0 {
		var err error
		rec.U, err = NewUser("martin", "password", "Martin")
		if err != nil {
			log.Fatal(err)
		}
		err = rec.U.GenKeys()
		if err != nil {
			log.Fatal(err)
		}
		err = InsertRecipient(rec)
		if err != nil {
			log.Fatal(err)
		}
	}

	order1.Index = 1
	order1.PanelSize = "100 1000 sq.ft homes each with their own private spaces for luxury"
	order1.TotalValue = 14000
	order1.Location = "India Basin, San Francisco"
	order1.MoneyRaised = 0
	order1.Metadata = "India Basin is an upcoming creative project based in San Francisco that seeks to invite innovators from all around to participate"
	order1.Live = false
	order1.INVAssetCode = ""
	order1.DEBAssetCode = ""
	order1.PBAssetCode = ""
	order1.DateInitiated = utils.Timestamp()
	order1.Years = 3
	order1.OrderRecipient = rec
	order1.Stage = 3
	err = InsertOrder(order1)
	if err != nil {
		return fmt.Errorf("Error inserting order into db")
	}

	order1.Index = 2
	order1.PanelSize = "180 1200 sq.ft homes in a high rise building 0.1mi from Kendall Square"
	order1.TotalValue = 30000
	order1.Location = "Kendall Square, Boston"
	order1.MoneyRaised = 0
	order1.Metadata = "Kendall Square is set in the heart of Cambridge and is a popular startup IT hub"
	order1.Live = false
	order1.INVAssetCode = ""
	order1.DEBAssetCode = ""
	order1.PBAssetCode = ""
	order1.DateInitiated = utils.Timestamp()
	order1.Years = 5
	order1.OrderRecipient = rec
	order1.Stage = 3
	err = InsertOrder(order1)
	if err != nil {
		return fmt.Errorf("Error inserting order into db")
	}

	order1.Index = 3
	order1.PanelSize = "260 1500 sq.ft homes set in a medieval cathedral style construction"
	order1.TotalValue = 40000
	order1.Location = "Trafalgar Square, London"
	order1.MoneyRaised = 0
	order1.Metadata = "Trafalgar Square is set in the heart of London's financial district, with big banks all over"
	order1.Live = false
	order1.INVAssetCode = ""
	order1.DEBAssetCode = ""
	order1.PBAssetCode = ""
	order1.DateInitiated = utils.Timestamp()
	order1.Years = 7
	order1.OrderRecipient = rec
	order1.Stage = 3
	err = InsertOrder(order1)
	if err != nil {
		return fmt.Errorf("Error inserting order into db")
	}

	var inv Investor
	allInvs, err := RetrieveAllInvestors()
	if err != nil {
		log.Fatal(err)
	}
	if len(allInvs) == 0 {
		var err error
		inv.U, err = NewUser("john", "password", "John")
		if err != nil {
			log.Fatal(err)
		}
		inv.VotingBalance = 100000
		err = inv.U.GenKeys()
		if err != nil {
			log.Fatal(err)
		}
		// TODO: this is being set as a constant now, but should be updated to check
		// the stablecoin and adjust accordingly.
		err = InsertInvestor(inv)
		if err != nil {
			log.Fatal(err)
		}
	}
	// NewOriginator(uname string, pwd string, Name string, Address string, Description string)
	newOriginator, err := NewOriginator("john", "password", "John Doe", "14 ABC Street London", "This is a sample originator")
	if err != nil {
		log.Fatal(err)
	}

	pc, err := newOriginator.OriginContract("100 16x24 panels on a solar rooftop", 14000, "Puerto Rico", 5, "ABC School in XYZ peninsula", 1) // 1 is the idnex for martin
	if err != nil {
		log.Fatal(err)
	}

	biddingOrder, err := RetrieveOrder(pc.O.Index)
	if err != nil {
		log.Fatal(err)
	}

	// Each contractor building off of this must reference the order index in their
	// proposed contract to enable searchability of the bucket. And each contractor
	// must build off of this in their proposed Contracts
	// Contractor stuff below
	contractor1, err := NewContractor("john", "password", "John Doe", "14 ABC Street London", "This is a sample contractor")
	if err != nil {
		log.Println(err)
	}

	_, err = contractor1.ProposeContract(pc.O.PanelSize, 28000, "Puerto Rico", 6, pc.O.Metadata+" we supply our own devs and provide insurance guarantee as well. Dual audit maintenance upto 1 year. Returns capped as per defaults", 1, biddingOrder.Index) // 1 for retrieving martin as the recipient
	if err != nil {
		log.Fatal(err)
	}

	// competing contractor details follow
	contractor2, err := NewContractor("sam", "password", "Samuel Jackson", "14 ABC Street London", "This is a competing contractor")
	if err != nil {
		log.Fatal(err)
	}

	_, err = contractor2.ProposeContract(pc.O.PanelSize, 35000, "Puerto Rico", 5, pc.O.Metadata+" free lifetime service, developers and insurance also provided", 1, biddingOrder.Index) // 1 for retrieving martin as the recipient
	if err != nil {
		log.Fatal(err)
	}

	_, err = NewOriginator("samuel", "password", "Samuel L. Jackson", "ABC Street, London", "I am an originator")
	if err != nil {
		log.Fatal(err)
	}

	_, err = RetrieveAllContractEntities("originator")
	if err != nil {
		log.Fatal(err)
	}
	_, err = RetrieveAllContractEntities("contractor")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
