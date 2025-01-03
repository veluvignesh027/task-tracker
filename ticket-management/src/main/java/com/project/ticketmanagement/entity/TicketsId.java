package com.project.ticketmanagement.entity;

import lombok.*;

@Data
@Getter
@Setter
public class TicketsId {
    private Integer ticketId;
    private String userStoryNo;

    public TicketsId() {
    }

    public TicketsId(Integer ticketId, String userStoryNo) {
        this.ticketId = ticketId;
        this.userStoryNo = userStoryNo;
    }

    public static TicketsId of(Integer ticketId, String userStoryNo) {
        return new TicketsId(ticketId, userStoryNo);
    }
}
