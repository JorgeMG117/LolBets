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
import com.example.lolbets.model.ActiveBets
import com.example.lolbets.viewmodel.ActiveBetsViewModel
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


class MyWebSocketListener(private val viewModel: FocusedGameViewModel, private val onBetSuccess: (ActiveBets) -> Unit) : WebSocketListener() {
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
                if(errorResponse.error == "success") {
                    viewModel.betSucceed()
                    println(viewModel.uiState.value.game.completed)
                    val activeBet = ActiveBets(viewModel.uiState.value.lastBet, viewModel.uiState.value.game.team1.code, viewModel.uiState.value.game.team2.code, viewModel.uiState.value.game.league.name, viewModel.uiState.value.game.completed)
                    onBetSuccess(activeBet)
                    //Añadir active bet

                }
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