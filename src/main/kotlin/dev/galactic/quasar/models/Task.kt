package dev.galactic.quasar.models

import java.time.LocalDateTime
import java.util.UUID

data class Task(
    val id: UUID = UUID.randomUUID(),
    val title: String,
    val description: String?,
    val dueDate: LocalDateTime?,
    val taskStatus: TaskStatus = TaskStatus.NOT_DONE,
)

enum class TaskStatus {
    NOT_DONE,
    DONE
}
