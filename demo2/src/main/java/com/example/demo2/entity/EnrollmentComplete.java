package com.example.demo2.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import lombok.Getter;
import lombok.NoArgsConstructor;

@Entity
@Getter
@NoArgsConstructor
public class EnrollmentComplete {
    @Id
    @GeneratedValue
    private Long id;

    @Column(name = "student_id")
    private Long studentId;

    @Column(name = "subject_id")
    private Long subjectId;

    public EnrollmentComplete(Long studentId, Long subjectId) {
        this.studentId = studentId;
        this.subjectId = subjectId;
    }
}
