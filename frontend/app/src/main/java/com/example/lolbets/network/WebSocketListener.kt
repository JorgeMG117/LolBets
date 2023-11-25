package com.example.lolbets.network

import com.example.lolbets.model.Bet
import com.example.lolbets.viewmodel.FocusedGameViewModel
import kotlinx.coroutines.flow.*
import kotlinx.serialization.Serializable
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import kotlinx.coroutines.flow.Flow
import okhttp3.Response
import okhttp3.WebSocket
import okhttp3.WebSocketListener
import android.util.Log
import kotlinx.serialization.SerialName
import okio.ByteString

@Serializable
sealed class WebsocketResponse

@Serializable
data class BetResponse(
    val id: Int,
    val bets1: Int,
    val bets2: Int,
) : WebsocketResponse()

@Serializable
data class ErrorResponse(
    @SerialName("error") val errorMessage: String
) : WebsocketResponse()


class MyWebSocketListener(private val viewModel: FocusedGameViewModel) : WebSocketListener() {
    override fun onOpen(webSocket: WebSocket, response: Response) {
        println("WebSocket connection opened")
        viewModel.setStatus(true)
        // You can send messages here using webSocket.send("Your message")
    }

    override fun onMessage(webSocket: WebSocket, text: String) {
        //{"id":8,"bets1":10,"bets2":15}
        //{"error":"success"}
        println("Received message: $text")
        //val message : Message = Json.decodeFromString(text)
        val result: WebsocketResponse = try {
            Json.decodeFromString(text)
        } catch (e: Exception) {
            // Handle parsing error
            println("Unable to parse the JSON string")
            return
        }

        when (result) {
            is BetResponse -> {
                println("Success: $result")
                viewModel.updateGameBets(result)
                // Access fields like result.id, result.bets1, result.bets2
            }
            is ErrorResponse -> {
                println("Error: $result")
                // Access the error message like result.errorMessage
            }
        }
        //viewModel.addMessage(Pair(false, text))

    }

    override fun onClosed(webSocket: WebSocket, code: Int, reason: String) {
        viewModel.setStatus(false)
        println("WebSocket closed: $code, $reason")
    }

    override fun onFailure(webSocket: WebSocket, t: Throwable, response: Response?) {
        println("WebSocket failure: ${t.message}")
    }
}