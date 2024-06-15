package com.batarilo.userservice.dto;


import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

import javax.validation.constraints.Min;
import javax.validation.constraints.NotNull;


@Builder
@AllArgsConstructor
@Getter
@Setter
public class RegisterUserDto {
    @NotNull
    private String email;
    @NotNull
    @Min(6)
    private String password;
    @NotNull
    private String fullName;
}
