package com.geo.LdapSearch.service;

import org.springframework.stereotype.Service;

@Service
public interface UserService {
    String getUser(String userName);
}
