package models

import (
	"fmt"
	"strings"
)

type Description interface {
}

type DescriptionStruct struct {
	SenderAccountNumber   string `json:"senderAccountNumber,omitempty"`
	RecieverAccountNumber string `json:"recieverAccountNumber,omitempty"`
	SenderName            string `json:"senderName,omitempty"`
	RecieverName          string `json:"recieverName,omitempty"`
	SenderAddress         string `json:"senderAddress,omitempty"`
	RecieverAddress       string `json:"recieverAddress,omitempty"`
	Description           string `json:"description,omitempty"`
}

func AddDescription1(data string, reciever bool) interface{} {
	if strings.Contains(data, ";") {
		desc := DescriptionStruct{}
		split := strings.Split(data, ";")
		switch len(split) {
		case 2:
			if reciever {
				desc.RecieverName = split[0]
				desc.RecieverAddress = split[1]
			} else {
				desc.SenderName = split[0]
				desc.SenderAddress = split[1]
			}
		case 3:
			if reciever {
				desc.RecieverName = split[0]
				desc.RecieverAddress = strings.Join(split[1:2], " ")
			} else {
				desc.SenderName = split[0]
				desc.SenderAddress = strings.Join(split[1:2], " ")
			}
		}
		return desc
	}
	return data
}

func AddDescription2(data string) interface{} {
	if strings.Contains(data, ";") {
		desc := DescriptionStruct{}
		split := strings.Split(data, ";")
		desc.Description = strings.Join(split[0:1], " ")
		return desc
	}
	return data
}

func AddDescription3(data string, reciever bool) interface{} {
	if strings.Contains(data, ";") {
		split := strings.Split(data, ";")
		fmt.Println(40, split, reciever)
		return data
	}
	return data
}
