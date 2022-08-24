module.exports = {
    "env": {
        "browser": true,
        "es2021": true,
        "amd": true,
        "node": true
    },
    "extends": [
        "eslint:recommended",
        "plugin:import/recommended",
        "plugin:import/typescript",
        "plugin:react/recommended",
        "plugin:@typescript-eslint/recommended"
    ],
    "overrides": [
    ],
    "parser": "@typescript-eslint/parser",
    "parserOptions": {
        "ecmaVersion": "latest",
        "sourceType": "module"
    },
    "plugins": [
        "react",
        "@typescript-eslint",
        "simple-import-sort"
    ],
    "rules": {
        "@typescript-eslint/no-empty-interface": 0,
        "@typescript-eslint/no-var-requires": 0,
        "import/no-unresolved": "error",
        "simple-import-sort/imports": "error",
        "simple-import-sort/exports": "error"
    }
}
