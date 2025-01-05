package com.batarilo.userservice.dto;

import lombok.Builder;
import lombok.Data;

@Builder
@Data
public class NewUserResponse {

    private Integer id;
    private String fullName;
    private String email;
}
