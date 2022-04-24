CREATE TABLE Game
(
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Actor_Nombre VARCHAR(60),
    Personaje_Nombre VARCHAR(60),
    Obra_Titulo VARCHAR(60),
    PRIMARY KEY (Obra_Titulo,Actor_Nombre,Personaje_Nombre),
    FOREIGN KEY (Obra_Titulo) REFERENCES Obra(Titulo),
    FOREIGN KEY (Actor_Nombre) REFERENCES Actor(Nombre),
    FOREIGN KEY (Personaje_Nombre) REFERENCES Personaje(Nombre)
);