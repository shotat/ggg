TODO:
  - [ ] DIの落とし所を考える

- application
  - input: usecaseのinput
  - output: usecaseのoutput
  - usecase: domainロジックの薄い調整タスク/application service
- domain
  - model
    - (root entity)
      - (entityA)
      - (entityB)
      - (VO)
      - ...
  - service
  - repository: 各集約に対応するrepository
- infrastructure
  - persistence
    - datastore: datastore用のrepository実装
    - ...
- interfaces
  - api
  - web
  - api_handler.go
- main.go

参考

- https://medium.com/@timakin/go%E3%81%AE%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E6%A7%8B%E6%88%90%E3%81%AE%E5%A4%B1%E6%95%97%E9%81%8D%E6%AD%B4%E3%81%A8%E7%8F%BE%E7%8A%B6%E7%A2%BA%E8%AA%8D-fc6a4369337
