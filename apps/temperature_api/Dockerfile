# Используем официальный образ OpenJDK 21
FROM eclipse-temurin:21-jdk-jammy

# Рабочая директория в контейнере
WORKDIR /app

# Копируем собранный JAR-файл в контейнер
COPY build/libs/temperature_api-0.0.1-SNAPSHOT.jar app.jar

# Открываем порт 8081
EXPOSE 8081

# Команда для запуска приложения
ENTRYPOINT ["java", "-jar", "app.jar"]