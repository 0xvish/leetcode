const fs = require("fs");
const path = require("path");

// Path to the problems file
const problemsFilePath = path.join(__dirname, "problems.json");

// Helper function to read the problems data
const readProblems = () => {
  if (fs.existsSync(problemsFilePath)) {
    return JSON.parse(fs.readFileSync(problemsFilePath, "utf-8"));
  } else {
    return {};
  }
};

// Helper function to write the problems data
const writeProblems = (problems) => {
  fs.writeFileSync(
    problemsFilePath,
    JSON.stringify(problems, null, 2),
    "utf-8"
  );
};

// Helper function to extract the problem number from the problem statement
const extractProblemNumber = (problemStatement) => {
  const match = problemStatement.match(/^(\d+)\./); // Matches the number at the beginning (e.g., "2270.")
  return match ? match[1] : null; // Returns the problem number (e.g., "2270")
};

// Function to organize problems by date and name
const organizeProblems = (date, problemStatement) => {
  const problems = readProblems();
  const problemNumber = extractProblemNumber(problemStatement); // Extract the problem number (e.g., "2270")
  if (!problemNumber) {
    console.log(
      "Problem statement is not in the correct format. Please include the problem number."
    );
    return;
  }

  const folderName = path.join(__dirname, `${date}_${problemNumber}`); // Use date and problem number as folder name

  // Create the folder if it doesn't exist
  if (!fs.existsSync(folderName)) {
    fs.mkdirSync(folderName);
  }

  // Add the problem to the problems object
  problems[date] = {
    name: problemStatement,
    folder: `${date}_${problemNumber}`,
  };

  writeProblems(problems);
  console.log(
    `Problem '${problemStatement}' organized under folder '${date}_${problemNumber}' for date '${date}'.`
  );
};

// Function to search for a problem by its name
const searchProblemByName = (problemName) => {
  const problems = readProblems();
  const results = Object.entries(problems).filter(([_, value]) =>
    value.name.toLowerCase().includes(problemName.toLowerCase())
  );

  if (results.length > 0) {
    console.log("Found matching problems:");
    results.forEach(([date, value]) =>
      console.log(`${date}: ${value.name} (Folder: ${value.folder})`)
    );
  } else {
    console.log(`No problems found with name containing '${problemName}'.`);
  }
};

// Function to search for a problem by its date
const searchProblemByDate = (date) => {
  const problems = readProblems();
  if (problems[date]) {
    console.log(
      `Problem on ${date}: ${problems[date].name} (Folder: ${problems[date].folder})`
    );
  } else {
    console.log(`No problem found for date '${date}'.`);
  }
};

// Function to search for a problem by its problem number
const searchProblemByNumber = (problemNumber) => {
  const problems = readProblems();
  const results = Object.entries(problems).filter(([_, value]) =>
    value.folder.includes(problemNumber)
  );

  if (results.length > 0) {
    console.log(`Found problems with problem number '${problemNumber}':`);
    results.forEach(([date, value]) =>
      console.log(`${date}: ${value.name} (Folder: ${value.folder})`)
    );
  } else {
    console.log(`No problems found with problem number '${problemNumber}'.`);
  }
};

// Function to list all problems
const listAllProblems = () => {
  const problems = readProblems();
  if (Object.keys(problems).length > 0) {
    console.log("List of all problems:");
    Object.entries(problems).forEach(([date, value]) => {
      console.log(`${date}: ${value.name} (Folder: ${value.folder})`);
    });
  } else {
    console.log("No problems found.");
  }
};

// Function to search for problems by a keyword
const searchProblemByKeyword = (keyword) => {
  const problems = readProblems();
  const results = Object.entries(problems).filter(([_, value]) =>
    value.name.toLowerCase().includes(keyword.toLowerCase())
  );

  if (results.length > 0) {
    console.log(`Found problems with keyword '${keyword}':`);
    results.forEach(([date, value]) =>
      console.log(`${date}: ${value.name} (Folder: ${value.folder})`)
    );
  } else {
    console.log(`No problems found with keyword '${keyword}'.`);
  }
};

// Command-line interface
const args = process.argv.slice(2);
const command = args[0];
const param1 = args[1];
const param2 = args[2];

switch (command) {
  case "organize":
    if (param1 && param2) {
      organizeProblems(param1, param2);
    } else {
      console.log("Usage: organize <date> <problem-statement>");
    }
    break;
  case "search-name":
    if (param1) {
      searchProblemByName(param1);
    } else {
      console.log("Usage: search-name <problem-name>");
    }
    break;
  case "search-date":
    if (param1) {
      searchProblemByDate(param1);
    } else {
      console.log("Usage: search-date <date>");
    }
    break;
  case "search-number":
    if (param1) {
      searchProblemByNumber(param1);
    } else {
      console.log("Usage: search-number <problem-number>");
    }
    break;
  case "list":
    listAllProblems();
    break;
  case "search-keyword":
    if (param1) {
      searchProblemByKeyword(param1);
    } else {
      console.log("Usage: search-keyword <keyword>");
    }
    break;
  default:
    console.log(
      "Invalid command. Usage: organize, search-name, search-date, list, search-keyword, search-number"
    );
    break;
}
