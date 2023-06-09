package com.example.demo.repository;

import com.example.demo.entity.EnrollmentPlan;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface EnrollmentPlanRepository extends JpaRepository<EnrollmentPlan, Long> {

    @Query("select ep from EnrollmentPlan ep where ep.studentId = :studentId")
    List<EnrollmentPlan> findByStudentId(@Param("studentId") Long studentId);
}
