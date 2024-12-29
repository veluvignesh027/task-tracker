package com.project.jira.api;

import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

import org.hibernate.sql.ast.tree.expression.NestedColumnReference;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;

import com.project.jira.entity.Tickets;
import com.project.jira.service.TicketService;

@RestController
public class TicketsAPI {
	
	@Autowired
	private TicketService ticketService;
	// api is used to get tickets by assignee
	@GetMapping("/ticket/assignee/{assignee}")
	public ResponseEntity<List<Tickets>> getTicketsByAssignee(@PathVariable String assignee) {
		Optional<List<Tickets>> ticketsList = ticketService.getTicketsByAssignee(assignee);
		if (ticketsList.isPresent()) {
			return new ResponseEntity<>(ticketsList.get(), HttpStatus.OK);
		}
		 else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}

	@GetMapping("/ticket/userStory/{userStory}")
	public ResponseEntity<List<Tickets>> getTicketsByUserStory(@PathVariable("userStory") String userStoryNo) {
		Optional<List<Tickets>> ticketsList = ticketService.getTicketsByUserStory(userStoryNo);
		if (ticketsList.isPresent()) {
			return new ResponseEntity<>(ticketsList.get(), HttpStatus.OK);
		}
		else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}
	@GetMapping("/ticket/{assignee}/{status}")
	public ResponseEntity<List<Tickets>> getTicketsByAssigneeAndTicketStatus(@PathVariable String assignee,@PathVariable("status") String ticketStatus) {
		Optional<List<Tickets>> ticketsList = ticketService.getTicketsByAssigneeAndTicketStatus(assignee,ticketStatus);
		if (ticketsList.isPresent()) {
			return new ResponseEntity<>(ticketsList.get(), HttpStatus.OK);
		}
		else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}

	@GetMapping("/ticketCount/assignee/{assignee}")
	public ResponseEntity<Integer> getCountOfTicketsByAssignee(@PathVariable String assignee) {
		Optional<List<Tickets>> ticketsList = ticketService.getTicketsByAssignee(assignee);
		//Long count =  Long.valueOf(ticketsList.map(List::size).orElse(0));
		Integer count =  ticketsList.map(List::size).orElse(0);
		if (count > 0) {
			return new ResponseEntity<>(count,HttpStatus.OK);
		}
		else {
			return new ResponseEntity<>(count,HttpStatus.NOT_FOUND);
		}
	}

	@GetMapping("/ticketCount/userStory/{userStory}")
	public ResponseEntity<Integer> getCountOfTicketsByUserStory(@PathVariable("userStory") String userStroyNo) {
		Optional<List<Tickets>> ticketsList = ticketService.getTicketsByUserStory(userStroyNo);
		//Long count =  Long.valueOf(ticketsList.map(List::size).orElse(0));
		Integer count =  ticketsList.map(List::size).orElse(0);
		if (count > 0) {
			return new ResponseEntity<>(count,HttpStatus.OK);
		}
		else {
			return new ResponseEntity<>(count,HttpStatus.NOT_FOUND);
		}
	}

	@GetMapping("/ticketCount/{assignee}/{status}")
	public ResponseEntity<Integer> getCountOfTicketsByAssigneeAndTicketStatus(@PathVariable String assignee,@PathVariable("status") String ticketStatus) {
		Optional<List<Tickets>> ticketsList = ticketService.getTicketsByAssigneeAndTicketStatus(assignee,ticketStatus);
		//Long count =  Long.valueOf(ticketsList.map(List::size).orElse(0));
		Integer count =  ticketsList.map(List::size).orElse(0);
		if (count > 0) {
			return new ResponseEntity<>(count,HttpStatus.OK);
		}
		else {
			return new ResponseEntity<>(count,HttpStatus.NOT_FOUND);
		}
	}

	@PostMapping("/addTicket")
	public ResponseEntity<?> addTicket(@RequestBody Tickets ticket){
		try{
			Tickets addedTicket = ticketService.addTicket(ticket);
			System.out.println(addedTicket);
			return new ResponseEntity<>(addedTicket,HttpStatus.CREATED);
		}
		catch (Exception e){
			return new ResponseEntity<>(e.getMessage(),HttpStatus.INTERNAL_SERVER_ERROR);
		}

	}

	@PutMapping("/updateTicket/{ticketId}")
	public ResponseEntity<?> updateTicket(@PathVariable Integer ticketId,@RequestBody Tickets ticket){
		long ticketCount = ticketService.countOfTicketAndUserStory(ticketId,ticket.getUserStoryNo());
		if (ticketCount > 0) {
			try{
				Tickets updatedTicket = ticketService.updateTicket(ticketId, ticket);
				System.out.println(updatedTicket);
				return new ResponseEntity<>(updatedTicket,HttpStatus.CREATED);
			}
			catch (Exception e){
				return new ResponseEntity<>(e.getMessage(),HttpStatus.INTERNAL_SERVER_ERROR);
			}
		}
		else{
			return new ResponseEntity<>("Record Not Exists",HttpStatus.NOT_FOUND);
		}
	}

	@DeleteMapping("/deleteTicket/userStory/{userStoryNo}")
	public ResponseEntity<?> deleteTicketByUserStry(@PathVariable String userStoryNo){
		long ticketCount = ticketService.countOfUserStory(userStoryNo);
		System.out.println(ticketCount);
		if (ticketCount > 0) {
			ticketService.deleteTicketByUserStoryNo(userStoryNo);
			return new ResponseEntity<>("Ticket Deleted", HttpStatus.OK);
		}
		else {
			return new ResponseEntity<>("Record Not Exists", HttpStatus.NOT_FOUND);
		}
	}

	@DeleteMapping(value = { "/deleteTicket/{userStoryNo}/{ticketId}",
							 "/deleteTicket/ticket/{ticketId}"
							})
	public ResponseEntity<?> deleteTicketByUserStryAndTicket(@PathVariable(required = false) String userStoryNo,@PathVariable Integer ticketId){
		System.out.println("inside controller");
		System.out.println("ticket id: "+ticketId);
		System.out.println("user story no: "+userStoryNo);
		if (userStoryNo == null)
			userStoryNo = "";
		System.out.println("user story no After assigned to null: "+userStoryNo);
		long ticketCount = ticketService.countOfTicketAndUserStory(ticketId,userStoryNo);
		System.out.println(ticketCount);
		if (ticketCount > 0) {
			ticketService.deleteTicketByIds(ticketId,userStoryNo);
			return new ResponseEntity<>("Ticket Deleted", HttpStatus.OK);
		}
		else {
			return new ResponseEntity<>("Record Not Exists", HttpStatus.NOT_FOUND);
		}
	}
}
