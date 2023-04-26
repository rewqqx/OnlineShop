package enums

// Будет переиспользованно при интеграции тестов внутрь проекта
enum class EButtonTypes(open val value: String) {

    // Tabs
    STATIC_TAB("static_tab_button"),
    CONTROL_TAB("control_tab_button"),
    PUZZLE_TAB("puzzle_tab_button"),
    OTHER_TAB("other_tab_button"),
    INTERACTIVITY_TAB("interactivity_tab_button"),
    DESIGN_TAB("design_tab_button"),

    // Filter TABS
    STATIC_FILTER_TAB("static_filter_tab_button"),
    CONTROL_FILTER_TAB("control_filter_tab_button"),
    PUZZLE_FILTER_TAB("puzzle_filter_tab_button"),
    OTHER_FILTER_TAB("other_filter_tab_button"),

    // Static Buttons
    CREATE_MARKER_BUTTON("create_marker_button"),
    CREATE_FRAME_BUTTON("create_frame_button"),
    CREATE_AREA_BUTTON("create_area_button"),
    CREATE_ARROW_BUTTON("create_arrow_button"),
    CREATE_LINE_BUTTON("create_line_button"),

    // Create Shared Buttons
    CREATE_IMAGE_BUTTON("create_image_button"),
    CREATE_RECT_BUTTON("create_rect_button"),

    // Create Control Button
    CREATE_BUTTON_BUTTON("create_button_button"),
    CREATE_INPUT_BUTTON("create_input_button"),
    CREATE_SELECTOR_BUTTON("create_selector_button"),

    // Create Puzzle Button
    CREATE_CUT_BUTTON("create_cut_button"),

    // Collapses
    ELEMENT_COLLAPSE("element_collapse_button"),
    EXPORT_COLLAPSE("export_collapse_button"),

    // Export Buttons
    DOWNLOAD_HTML_BUTTON("download_html_button"),
    PREVIEW_BUTTON("preview_button"),
    SHARE_BUTTON("share_button"),
    INTEGRATE_BUTTON("share_button"),

    // Header Buttons
    FILE_BUTTON("header_file_button"),
    LOAD_JSON_BUTTON("header_load_json_button"),
    UPLOAD_JSON_BUTTON("header_upload_json_button"),
    UNDO_BUTTON("header_undo_button"),
    REDO_BUTTON("header_redo_button"),
    COPY_BUTTON("header_copy_button"),
    DELETE_BUTTON("header_delete_button"),
    STYLE_BUTTON("header_style_button"),

    // Interactivity Inputs
    HINT_INPUT("hint_input"),
    PLACEHOLDER_INPUT("placeholder_input"),
    HINT_ACTIVE_SELECTOR("hint_active_selector"),
    ELEMENT_SHOWING_SELECTOR("hint_active_selector"),
    CORRECT_STATE_BUTTON("correct_state_button"),

    // Design Buttons
    RESOLUTION_A_BUTTON("resolution_a_button"),
    RESOLUTION_B_BUTTON("resolution_b_button"),
    RESOLUTION_C_BUTTON("resolution_c_button"),


}