package com.batarilo.userservice.service;

import com.batarilo.userservice.dto.LoginUserDto;
import com.batarilo.userservice.dto.NewUserResponse;
import com.batarilo.userservice.dto.RegisterUserDto;
import com.batarilo.userservice.model.User;
import com.batarilo.userservice.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class AuthenticationService
{

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final AuthenticationManager authenticationManager;

    public NewUserResponse signup(RegisterUserDto input)
    {
        User user = User.builder()
            .email(input.getEmail())
            .password(passwordEncoder.encode(input.getPassword()))
            .fullName(input.getFullName())
            .build();

     User resultUser = userRepository.save(user);

     return NewUserResponse.builder()
             .id(resultUser.getId())
             .email(resultUser.getEmail())
             .fullName(resultUser.getFullName())
             .build();
    }

    public User authenticate(LoginUserDto input)
    {
        authenticationManager.authenticate(
            new UsernamePasswordAuthenticationToken(
                input.getEmail(),
                input.getPassword()
            )
        );

        return userRepository.findByEmail(input.getEmail())
            .orElseThrow(() -> new BadCredentialsException("Bad credentials."));
    }

    public Boolean isEmailInUse(String email)
    {
        return userRepository
            .findByEmail(email)
            .isPresent();
    }
}
