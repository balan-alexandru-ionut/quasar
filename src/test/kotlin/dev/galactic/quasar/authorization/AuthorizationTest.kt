package dev.galactic.quasar.authorization

import dev.galactic.quasar.util.fakeJwt
import io.github.serpro69.kfaker.faker
import io.kotest.matchers.comparables.shouldBeGreaterThan
import io.kotest.matchers.shouldBe
import io.quarkus.test.junit.QuarkusTest
import io.smallrye.jwt.auth.principal.JWTParser
import jakarta.inject.Inject
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import java.time.Duration
import java.time.Instant
import java.util.UUID

@QuarkusTest
class AuthorizationTest {
    @Inject
    lateinit var jwtParser: JWTParser

    private val faker = faker {  }

    @Test
    @DisplayName("given a jwt token created by using the dev keys, when trying to verify it, then the operation should succeed")
    fun jwtCreateAndVerify() {
        val userId = UUID.randomUUID().toString()
        val userEmail = faker.internet.email()

        val jwt = fakeJwt(
            duration = Duration.ofMinutes(5),
            subject = userId,
            extraClaims = listOf(Pair("userEmail", userEmail))
        )

        val parsedToken = jwtParser.parse(jwt)

        parsedToken.issuer shouldBe "Quasar"
        parsedToken.getClaim<String>("userEmail") shouldBe userEmail
        parsedToken.expirationTime shouldBeGreaterThan Instant.now().epochSecond
    }
}