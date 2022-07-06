 inline_policy {
          + name   = "task-role-custom"
          + policy = jsonencode(
                {
                  + Statement = [
                      + {
                          + Action   = [
                              + "sqs:SendMessage",
                            ]
                          + Effect   = "Allow"
                          + Resource = [
                              + "arn:aws:sqs:ap-northeast-2:067084483518:UserAccountPointJobQueue",
                            ]
                        },
                      + {
                          + Action   = [
                              + "s3:PutObject",
                              + "s3:PutObjectAcl",
                              + "s3:DeleteObject",
                            ]
                          + Effect   = "Allow"
                          + Resource = [
                              + "arn:aws:s3:::hub.data.alpha.zigzag.kr/*",
                              + "arn:aws:s3:::res.s.alpha.zigzag.kr/*",
                              + "arn:aws:s3:::storage.alpha.zigzag.kr/*",
                              + "arn:aws:s3:::content.alpha.zigzag.kr/*",
                              + "arn:aws:s3:::marketing.s.alpha.zigzag.kr/*",
                            ]
                        },
                      + {
                          + Action   = [
                              + "s3:GetObject",
                              + "s3:GetObjectAcl",
                              + "s3:ListBucket",
                            ]
                          + Effect   = "Allow"
                          + Resource = [
                              + "arn:aws:s3:::hub.data.alpha.zigzag.kr/*",
                              + "arn:aws:s3:::res.s.alpha.zigzag.kr/*",
                              + "arn:aws:s3:::storage.alpha.zigzag.kr/*",
                              + "arn:aws:s3:::content.alpha.zigzag.kr/*",
                              + "arn:aws:s3:::marketing.s.alpha.zigzag.kr/*",
                            ]
                        },
                      + {
                          + Action   = [
                              + "sts:AssumeRole",
                            ]
                          + Effect   = "Allow"
                          + Resource = [
                              + "arn:aws:iam::528112856704:role/AthenaDevBatchReadRole",
                            ]
                        },
                      + {
                          + Action   = [
                              + "ssm:GetParameter",
                            ]
                          + Effect   = "Allow"
                          + Resource = [
                              + "arn:aws:ssm:ap-northeast-2:067084483518:parameter/github-token",
                            ]
                        },
                    ]
                  + Version   = "2012-10-17"
                }
            )
