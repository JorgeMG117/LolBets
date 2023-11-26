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
    val error: String
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
        try {
            // Try to parse as a SuccessResponse
            val betResponse: BetResponse = Json.decodeFromString(text)
            viewModel.updateGameBets(betResponse)
            println("Received message: $betResponse")
        } catch (e: Exception) {
            try {
                // Try to parse as an ErrorResponse
                val errorResponse: ErrorResponse = Json.decodeFromString(text)
                println("Received message: $errorResponse")
                //TODO Show a text to say if the bet was process successfully
            } catch (e: Exception) {
                println("Unable to parse the JSON string")
            }
        }
    }

    override fun onClosed(webSocket: WebSocket, code: Int, reason: String) {
        viewModel.setStatus(false)
        println("WebSocket closed: $code, $reason")
    }

    override fun onFailure(webSocket: WebSocket, t: Throwable, response: Response?) {
        println("WebSocket failure: ${t.message}")
    }
}