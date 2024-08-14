package com.geo.LdapSearch.controller;

import com.geo.LdapSearch.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class UserSearchController {
    @Autowired
    private UserService service;
    @GetMapping("/user/{userId}")
    public ResponseEntity<String> getUser(@PathVariable("userId") String name){
        return ResponseEntity.ok(service.getUser(name));
    }

}
