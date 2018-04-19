class CaluclatePrice
  attr_reader :distance
  attr_reader :time

  # distance - дистанция в метрах
  # time - время в секундах

  # тариф
  # 250 ₽ подача + 20 ₽ за мин + 20 ₽ за км,
  # минимальная стоимость 500 ₽

  Supply = 250
  MinPrice = 500
  PerMinute = 20
  PerKm = 20

  def initialize(distance, time)
    @distance = distance
    @time = time
  end

  def run
    distance_price = (distance / 1000.0) * PerKm
    time_price = (time / 60.0) * PerMinute

    price = (Supply + distance_price + time_price).round

    price < MinPrice ? MinPrice : price
  end
end
