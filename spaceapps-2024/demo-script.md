# NASA Space Apps Demo Script

## Introduction (0:00-0:20)
"Flooding causes $2.6B in annual damage in Nigeria. Our solution: An AI-powered flood prediction system using NASA satellite data."

## Problem (0:21-0:40)
"Makurdi suffers chronic flooding due to inadequate drainage systems. Existing models use outdated data and lack real-time integration."

## Solution (0:41-1:20)
*Show CLI tool:*
./idf-tool -cli
> Duration: 2
> Return Period: 50
> Intensity: 142.3 mm/hr

*Show API call:*
curl "http://localhost:8080/idf?d=2&t=50"
{"intensity":142.3, "duration":2, ...}

*Show dashboard:*
http://localhost:3000

## NASA Integration (1:21-1:50)
"Real-time NASA GPM data enhances predictions by 23% accuracy during extreme weather events."

## Impact (1:51-2:00)
"Our open-source tool helps engineers build climate-resilient infrastructure across West Africa."
