package com.project.ticketmanagement.entity;

import jakarta.persistence.*;
import org.hibernate.annotations.GenericGenerator;

import lombok.Data;
import lombok.Getter;
import lombok.Setter;

@Entity
@IdClass(TicketsId.class)
@Data
@Getter @Setter
public class Tickets extends BaseEntity{
	
	@Id
	//@GeneratedValue(strategy = GenerationType.IDENTITY)
	private Integer ticketId;
	@Id
	private String userStoryNo;
	private String ticketName;
	private String ticketDesc;
	private String assignee;
	private String ticketStatus;

	public Integer getTicketId() {
		return ticketId;
	}
	public void setTicketId(Integer ticketId) {
		this.ticketId = ticketId;
	}

	public String getUserStoryNo() {
		return userStoryNo;
	}

	public void setUserStoryNo(String userStoryNo) {
		this.userStoryNo = userStoryNo;
	}

	public String getTicketName() {
		return ticketName;
	}

	public void setTicketName(String ticketName) {
		this.ticketName = ticketName;
	}

	public String getTicketDesc() {
		return ticketDesc;
	}

	public void setTicketDesc(String ticketDesc) {
		this.ticketDesc = ticketDesc;
	}

	public String getAssignee() {
		return assignee;
	}

	public void setAssignee(String assignee) {
		this.assignee = assignee;
	}

	public String getTicketStatus() {
		return ticketStatus;
	}

	public void setTicketStatus(String ticketStatus) {
		this.ticketStatus = ticketStatus;
	}

	@Override
	public String toString() {
		return "Tickets{" +
				"ticketId=" + ticketId +
				", userStoryNo=" + userStoryNo +
				", ticketName='" + ticketName + '\'' +
				", ticketDesc='" + ticketDesc + '\'' +
				", assignee='" + assignee + '\'' +
				", ticketStatus='" + ticketStatus + '\'' +
				'}';
	}
}
