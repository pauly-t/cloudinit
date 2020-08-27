package iso9660

var iso9660DefaultID = []int8{67, 68, 49, 48, 48} // CD100

type iso9660VolumeDescriptor struct {
	_type   int8       // (1, 1)
	id      [5]int8    // (2, 6)
	version int8       // (7, 7)
	data    [2041]int8 // (8, 2048)
}

var iso9660PrimaryDescriptor struct {
	_type                  int8    // (1,    1)
	id                     [5]int8 // (2,    6)
	version                int8    // (7,    7)
	unused1                int8    // (8,    8)
	system_id                      // (9,   40)
	volume_id                      // (41,  72)
	unused2                        // (73,  80)
	volume_space_size              // (81,  88)
	escape_sequences               // (89, 120)
	volume_set_size                // (121, 124)
	volume_sequence_number         // (125, 128)
	logical_block_size             // (129, 132)
	path_table_size                // (133, 140)
	type_l_path_table              // (141, 144)
	opt_type_l_path_table          // (145, 148)
	type_m_path_table              // (149, 152)
	opt_type_m_path_table          // (153, 156)
	root_directory_record          // (157, 190)
	volume_set_id                  // (191, 318)
	publisher_id                   // (319, 446)
	preparer_id                    // (447, 574)
	application_id                 // (575, 702)
	copyright_file_id              // (703, 739)
	abstract_file_id               // (740, 776)
	bibliographic_file_id          // (777, 813)
	creation_date                  // (814, 830)
	modification_date              // (831, 847)
	expiration_date                // (848, 864)
	effective_date                 // (865, 881)
	file_structure_version         // (882, 882)
	unused4                        // (883, 883)
	application_data               // (884, 1395)
	unused5                        // (1396, 2048)
}

// struct iso_path_table {
// unsigned char  name_len[2];	/* 721 */
// char extent[4];			/* 731 */
// char  parent[2];		/* 721 */
// char name[1];
// };

// struct iso_directory_record {
// unsigned char length		[ISODCL(1,  1)];  /* 711 */
// char ext_attr_length		[ISODCL(2,  2)];  /* 711 */
// char extent			[ISODCL(3,  10)]; /* 733 */
// char size			[ISODCL(11, 18)]; /* 733 */
// char date			[ISODCL(19, 25)]; /* 7 by 711 */
// unsigned char flags		[ISODCL(26, 26)];
// char file_unit_size		[ISODCL(27, 27)]; /* 711 */
// char interleave			[ISODCL(28, 28)]; /* 711 */
// char volume_sequence_number	[ISODCL(29, 32)]; /* 723 */
// unsigned char name_len		[ISODCL(33, 33)]; /* 711 */
// char name			[MAX_ISONAME+1]; /* Not really, but we need something here */
// };
