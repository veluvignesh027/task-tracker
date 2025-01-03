package com.project.ticketmanagement.repository;

import java.util.List;
import java.util.Optional;

import com.project.ticketmanagement.entity.TicketsId;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;
import com.project.ticketmanagement.entity.Tickets;

@Repository
public interface TicketRepository extends CrudRepository<Tickets, TicketsId>{

	 Optional<Tickets> findTopByUserStoryNoOrderByTicketIdDesc(String userStoryNo);
	 long countByTicketIdAndUserStoryNo(Integer ticketId, String userStoryNo);
	 long countByUserStoryNo(String userStoryNo);
	 void deleteByUserStoryNo(String userStoryNo);
	 List<Tickets> findByAssigneeOrderByUserStoryNoAscTicketIdAsc(String assignee);
	 List<Tickets> findByAssigneeIgnoreCaseAndTicketStatusIgnoreCaseOrderByUserStoryNoAscTicketIdAsc(String assignee,String ticketStatus);
	 List<Tickets> findByUserStoryNoOrderByTicketId(String userStoryNo);
}
