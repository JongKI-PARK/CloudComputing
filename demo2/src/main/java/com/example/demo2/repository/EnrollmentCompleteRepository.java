package com.example.demo2.repository;

import com.example.demo2.entity.EnrollmentComplete;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface EnrollmentCompleteRepository extends JpaRepository<EnrollmentComplete, Long> {
    @Query("select ec from EnrollmentComplete ec where ec.subjectId = :subjectId")
    Optional<EnrollmentComplete> findBySubjectId(Long subjectId);
}
