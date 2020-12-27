require "csv"

name_to_cost = Hash.new(0)

CSV.foreach(ARGV[0], headers: true) do |row|
  name_to_cost[row["name"]] += row["cost"].to_f
end

name_to_cost.sort_by {|k, _| k }.each do |name, cost|
  printf "%s\t%.3f\n", name, cost
end
