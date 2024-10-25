package dev.galactic.quasar.models

import jakarta.validation.constraints.Email
import java.util.UUID

data class Person(
    val id: UUID = UUID.randomUUID(),

    // Authentication
    @Email
    val email: String,
    val password: String,

    // Authorization
    val accessToken: String,
    val refreshToken: String,

    // Person Data
    val firstName: String,
    val lastName: String,
    val tasks: Set<Task>
)