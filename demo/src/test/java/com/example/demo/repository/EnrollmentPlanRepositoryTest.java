package com.example.demo.repository;

import com.example.demo.entity.EnrollmentPlan;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.jdbc.AutoConfigureTestDatabase;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;

import java.util.List;

import static org.assertj.core.api.Assertions.*;

@DataJpaTest
@AutoConfigureTestDatabase
class EnrollmentPlanRepositoryTest {
    @Autowired
    EnrollmentPlanRepository enrollmentPlanRepository;

    @Test
    @DisplayName("저장 확인")
    void saveTest() {
        EnrollmentPlan enrollmentPlan = new EnrollmentPlan(1L, 2L);

        EnrollmentPlan savedEnrollmentPlan = enrollmentPlanRepository.save(enrollmentPlan);

        List<EnrollmentPlan> results = enrollmentPlanRepository.findAll();
        assertThat(results.size()).isEqualTo(1);
        EnrollmentPlan result = results.get(0);
        assertThat(result.getStudentId()).isEqualTo(enrollmentPlan.getStudentId());
        assertThat(result.getSubjectId()).isEqualTo(enrollmentPlan.getSubjectId());
    }

    @Test
    @DisplayName("학생 id로 희망과목 조회")
    void findAllByStudentId() {
        for (int i = 0; i < 5; i++) {
            enrollmentPlanRepository.save(new EnrollmentPlan(1L, (long) (i + 1)));
            enrollmentPlanRepository.save(new EnrollmentPlan(2L, (long) (i + 1)));
        }

        List<EnrollmentPlan> results = enrollmentPlanRepository.findByStudentId(1L);

        assertThat(results.size()).isEqualTo(5);
        for (int i = 0; i < 5; i++)
            assertThat(results.get(i).getSubjectId()).isEqualTo(i + 1);
    }
}