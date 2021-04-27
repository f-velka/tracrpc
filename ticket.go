package tracrpc

import (
	"errors"
)

const (
	ticket_query                string = "ticket.query"
	ticket_getRecentChanges     string = "ticket.getRecentChanges"
	ticket_getAvailableActionsy string = "ticket.getAvailableActions"
	ticket_getActions           string = "ticket.getActions"
	ticket_get                  string = "ticket.get"
	ticket_create               string = "ticket.create"
	ticket_update               string = "ticket.update"
	ticket_delete               string = "ticket.delete"
	ticket_change_log           string = "ticket.changeLog"
	ticket_list_attachments     string = "ticket.listAttachments"
	ticket_get_attachment       string = "ticket.getAttachment"
	ticket_put_attachment       string = "ticket.putAttachment"
	ticket_delete_attachment    string = "ticket.deleteAttachment"
	ticket_get_ticket_fields    string = "ticket.getTicketFields"
)

// TicketService represents ticket API service.
type TicketService struct {
	rpc       RpcClient
	Component *TicketComponentService
}

// newTicketService creates new TicketService instance.
func newTicketService(rpc RpcClient) (*TicketService, error) {
	if rpc == nil {
		return nil, errors.New("rpc client cannot be nil")
	}

	component, err := newTicketComponentService(rpc)
	if err != nil {
		return nil, err
	}

	return &TicketService{
		rpc:       rpc,
		Component: component,
	}, nil
}

// Query calls ticket.query.
func (t *TicketService) Query(qstr) (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) GetRecentChanges() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) GetAvailableActions() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) GetActions() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) Get() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) Create() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) Update() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) Delete() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) ChangeLog() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) ListAttachments() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) GetAttachment() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) PutAttachment() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) DeleteAttachment() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

//  calls ticket..
func (t *TicketService) GetTicketFields() (string, error) {
	args := packArgs()
	var reply string
	if err := t.rpc.Call(ticket_query, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}
