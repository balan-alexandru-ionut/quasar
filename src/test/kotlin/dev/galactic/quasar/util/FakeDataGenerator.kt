package dev.galactic.quasar.util

import io.smallrye.jwt.build.Jwt
import java.time.Duration
import java.util.UUID

fun fakeJwt(
    duration: Duration,
    subject: String = UUID.randomUUID().toString(),
    extraClaims: List<Pair<String, String>>? = null
): String {
    return Jwt
        .subject(subject)
        .issuer("Quasar")
        .expiresIn(duration)
        .also { claimsSet ->
            extraClaims?.forEach {
                claimsSet.claim(it.first, it.second)
            }
        }
        .innerSign()
        .encrypt()
}