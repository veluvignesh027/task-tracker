package com.project.ticketmanagement.api;

import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

import org.hibernate.sql.ast.tree.expression.NestedColumnReference;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;

import com.project.ticketmanagement.entity.Tickets;
import com.project.ticketmanagement.service.TicketService;

@RestController
public class TicketsAPI {
	
	@Autowired
	private TicketService ticketService;
	// api is used to get tickets by assignee
	@GetMapping("/ticket/userStory/{userStory}")
	public ResponseEntity<?> getListOrCountOfTicketsByUserStory(@PathVariable("userStory") String userStoryNo,@RequestParam String query ) {
		Optional<List<Tickets>> ticketsList = ticketService.getTicketsByUserStory(userStoryNo);
		if (query.equals("records")){
			if (ticketsList.isPresent()) {
				return new ResponseEntity<>(ticketsList.get(), HttpStatus.OK);
			}
			else {
				return new ResponseEntity<>(HttpStatus.NOT_FOUND);
			}
		} else if (query.equals("count")) {
			Integer count =  ticketsList.map(List::size).orElse(0);
			if (count > 0) {
				return new ResponseEntity<>(count,HttpStatus.OK);
			}
			else {
				return new ResponseEntity<>(count,HttpStatus.NOT_FOUND);
			}
		}
		else {
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}
	}

	@GetMapping("/ticket/assignee/{assignee}")
	public ResponseEntity<?> getListOrCountOfTicketsByAssignee(@PathVariable String assignee,@RequestParam String query) {
		Optional<List<Tickets>> ticketsList = ticketService.getTicketsByAssignee(assignee);
		if (query.equals("records")){
			if (ticketsList.isPresent()) {
				return new ResponseEntity<>(ticketsList.get(), HttpStatus.OK);
			}
			else {
				return new ResponseEntity<>(HttpStatus.NOT_FOUND);
			}
		} else if (query.equals("count")) {
			Integer count =  ticketsList.map(List::size).orElse(0);
			if (count > 0) {
				return new ResponseEntity<>(count,HttpStatus.OK);
			}
			else {
				return new ResponseEntity<>(count,HttpStatus.NOT_FOUND);
			}
		}
		else {
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}
	}

	@GetMapping("/ticket/{assignee}")
	public ResponseEntity<?> getListOrCountOfTicketsByAssigneeAndTicketStatus(@PathVariable String assignee,@RequestParam String query,@RequestParam("status") String ticketStatus) {
		Optional<List<Tickets>> ticketsList = ticketService.getTicketsByAssigneeAndTicketStatus(assignee,ticketStatus);
		if (query.equals("records")){
			if (ticketsList.isPresent()) {
				return new ResponseEntity<>(ticketsList.get(), HttpStatus.OK);
			}
			else {
				return new ResponseEntity<>(HttpStatus.NOT_FOUND);
			}
		} else if (query.equals("count")) {
			Integer count =  ticketsList.map(List::size).orElse(0);
			if (count > 0) {
				return new ResponseEntity<>(count,HttpStatus.OK);
			}
			else {
				return new ResponseEntity<>(count,HttpStatus.NOT_FOUND);
			}
		}
		else {
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
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
			return new ResponseEntity<>("Record Not Exists",HttpStatus.NO_CONTENT);
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
			return new ResponseEntity<>("Record Not Exists", HttpStatus.NO_CONTENT);
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
			return new ResponseEntity<>("Record Not Exists", HttpStatus.NO_CONTENT);
		}
	}
}
