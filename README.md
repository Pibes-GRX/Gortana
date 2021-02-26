# Gortana: Asistente Virtual en Telegram

## Integrantes
 - [José Pablo Márquez Megías](https://www.linkedin.com/in/jose-pablo-m%C3%A1rquez-meg%C3%ADas-3942241a3/)
 - [Miguel Ángel Martín Rodríguez](https://www.linkedin.com/in/miguel-%C3%A1ngel-mart%C3%ADn-rodr%C3%ADguez-2893571a3/)
 - [Miquel Reyes Castells](https://www.linkedin.com/in/miquel-reyes/)

## Descripción del proyecto
Nuestro objetivo es crear un bot de Telegram que haga un análisis estadístico de los mensajes en un chat, siendo capaz de extraer quien es la persona que más habla, cuáles son las palabras más usadas, además de hacer de agente conversacional (obviamente estará muy limitado pues se usará una batería de palabras y respuestas, en lugar de una Inteligencia Artificial como tal).

## Solución
Para abordar este problema vamos a desarrollar un hook (es decir, un programa que reaccionará ante las peticiones realizadas a través de una funcionalidad a la que esté suscrita) usando la API de Telegram para Go [Telebot](https://github.com/tucnak/telebot), en conjunto con un diccionario de respuestas.

## Tecnologías que usaremos
- Logger: [Paquete log.go](https://golang.org/pkg/log/)
- Configuración Remota: Próximamente
- BBDD: Próximamente
