package com.project.jira.service;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

import com.project.jira.entity.TicketsId;
import jakarta.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.project.jira.entity.Tickets;
import com.project.jira.repository.TicketRepository;

@Service
public class TicketService {
	
	@Autowired
	private TicketRepository ticketRepo;

	public long countOfTicketAndUserStory(Integer ticketId,String userStoryNo) {
		return ticketRepo.countByTicketIdAndUserStoryNo(ticketId,userStoryNo);
	}

	public long countOfUserStory(String userStoryNo) {
		return ticketRepo.countByUserStoryNo(userStoryNo);
	}

	public Tickets addTicket(Tickets ticket) {
		int existingTicket = ticketRepo.findTopByUserStoryNoOrderByTicketIdDesc(ticket.getUserStoryNo())
				.map(Tickets::getTicketId)
				.orElse(0);
		if (existingTicket >0){
			ticket.setTicketId(existingTicket+1);
		}
		else{
			ticket.setTicketId(1);
		}
		ticket.setCreatedAt(LocalDateTime.now());
		ticket.setCreatedBy("Saran");
		return ticketRepo.save(ticket);
	}

	public Tickets updateTicket(Integer ticketId, Tickets ticket) {
		ticket.setTicketId(ticketId);
		ticket.setLastModifiedBy("Unknown");
		ticket.setLastModifiedAt(LocalDateTime.now());
		return ticketRepo.save(ticket);
	}

	public void deleteTicketByIds(Integer ticketId, String userStoryNo) {
		TicketsId ticketsId = TicketsId.of(ticketId,userStoryNo);
		ticketRepo.deleteById(ticketsId);
	}

	@Transactional
	public void deleteTicketByUserStoryNo(String userStoryNo) {
		ticketRepo.deleteByUserStoryNo(userStoryNo);
	}
	public Optional<List<Tickets>> getTicketsByAssignee(String assignee) {
		List<Tickets> ticketsList = ticketRepo.findByAssigneeOrderByUserStoryNoAscTicketIdAsc(assignee);
		return ticketsList.isEmpty() ? Optional.empty() : Optional.of(ticketsList);
	}

	public Optional<List<Tickets>> getTicketsByAssigneeAndTicketStatus(String assignee,String ticketStatus) {
		List<Tickets> ticketsList = ticketRepo.findByAssigneeIgnoreCaseAndTicketStatusIgnoreCaseOrderByUserStoryNoAscTicketIdAsc(assignee,ticketStatus);
		return ticketsList.isEmpty() ? Optional.empty() : Optional.of(ticketsList);
	}

	public Optional<List<Tickets>> getTicketsByUserStory(String userStoryNo) {
		List<Tickets> ticketsList = ticketRepo.findByUserStoryNoOrderByTicketId(userStoryNo);
		return ticketsList.isEmpty() ? Optional.empty() : Optional.of(ticketsList);
	}
}
