package com.example.demo2.repository;

import com.example.demo2.entity.EnrollmentOrder;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface EnrollmentOrderRepository extends JpaRepository<EnrollmentOrder, Long> {
}
