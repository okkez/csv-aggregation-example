import org.apache.commons.csv.CSVFormat
import java.io.File

fun main(args: Array<String>) {
    val nameToCost = mutableMapOf<String, Float>()
    File(args[0]).bufferedReader().use {
        CSVFormat.DEFAULT.withHeader("id", "name", "description", "cost").withFirstRecordAsHeader().parse(it).forEach {
            nameToCost.put(it.get("name"), nameToCost.getOrDefault(it.get("name"), 0.0F) + it.get("cost").toFloat())
        }
    }
    for (entry in nameToCost.toSortedMap()) {
        println("%s\t%.3f".format(entry.key, entry.value))
    }
}