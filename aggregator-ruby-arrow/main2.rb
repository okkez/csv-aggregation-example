require "arrow"

schema = Arrow::Schema.new([
  Arrow::Field.new("id", :string),
  Arrow::Field.new("name", :string),
  Arrow::Field.new("description", :string),
  Arrow::Field.new("string", :float),
])

name_to_cost = Hash.new(0)
table = Arrow::Table.load(ARGV[0], schema: schema)
table = table.select_columns do |column|
  %w(name cost).include?(column.name)
end
g = table.group("name")
result = g.sum
result.each_record_batch do |batch|
  batch.each do |row|
    name_to_cost[row["name"]] += row["cost"]
  end
end


name_to_cost.sort_by {|k, _| k }.each do |name, cost|
  printf("%s\t%.3f\n", name, cost)
end

__END__
