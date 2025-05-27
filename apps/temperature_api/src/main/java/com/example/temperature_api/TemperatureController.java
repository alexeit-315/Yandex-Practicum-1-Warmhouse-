package com.example.temperature_api;

import org.springframework.web.bind.annotation.*;
import java.util.concurrent.ThreadLocalRandom;
import java.time.Instant;

@RestController
public class TemperatureController {

    @GetMapping("/temperature")
    public TemperatureResponse getTemperature(@RequestParam String location) {
        String sensorId, status, description;
        switch (location) {
            case "Living Room":
                sensorId = "1";
                status = "active";
                description = "Living Room sensor";
                break;
            case "Bedroom":
                sensorId = "2";
                status = "active";
                description = "Bedroom sensor";
                break;
            case "Kitchen":
                sensorId = "3";
                status = "active";
                description = "Kitchen sensor";
                break;
            default:
                location = "Unknown";
                status = "inactive";
                description = "Unknown sensor";
                sensorId = "0";
        }

        return new TemperatureResponse(
                generateTemperatureResponse(),
                "C",
                Instant.now(),
                location,
                status,
                sensorId,
                "temperature",
                description
        );
    }

    @GetMapping("/temperature/{sensorId}")
    public TemperatureResponse getTemperatureById(@PathVariable String sensorId) {
        String location, status, description;
        switch (sensorId) {
            case "1":
                location = "Living Room";
                status = "active";
                description = "Living Room sensor";
                break;
            case "2":
                location = "Bedroom";
                status = "active";
                description = "Bedroom sensor";
                break;
            case "3":
                location = "Kitchen";
                status = "active";
                description = "Kitchen sensor";
                break;
            default:
                location = "Unknown";
                description = "Unknown sensor";
                status = "inactive";
        }

        return new TemperatureResponse(
                generateTemperatureResponse(),
                "C",
                Instant.now(),
                location,
                status,
                sensorId,
                "temperature",
                description
        );
    }

    private double generateTemperatureResponse() {
        return Math.round((10 + ThreadLocalRandom.current().nextDouble() * 20) * 10) / 10.0;
    }

    record TemperatureResponse(
            double value,
            String unit,
            Instant timestamp,
            String location,
            String status,
            String sensor_id,
            String sensor_type,
            String description
    ) {}
}